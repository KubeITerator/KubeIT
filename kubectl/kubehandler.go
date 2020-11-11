package kubectl

import (
	"bytes"
	"flag"
	"fmt"
	argov1alpha "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	wfv1 "github.com/argoproj/argo/pkg/client/clientset/versioned/typed/workflow/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	k8sYaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type KubeHandler struct {
	client wfv1.WorkflowInterface
}

func (kube *KubeHandler) StartClient(namespace string) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	//config, err := rest.InClusterConfig()
	//if err != nil {
	//	panic(err)
	//}
	clientset, err := wfv1.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	kube.client = clientset.Workflows(namespace)

}

func (kube *KubeHandler) StartJob(yaml string) error {

	wf, err := kube.ValidateYaml(yaml)
	if err != nil {
		fmt.Println("Error in validating Yaml")
		fmt.Println("Yaml:")
		fmt.Print(yaml)
		return err
	}
	// Create Deployment
	fmt.Println("Creating job...")
	result, err := kube.client.Create(wf)
	if err != nil {
		fmt.Println("Error in creating Job")
		fmt.Println("Yaml:")
		fmt.Print(yaml)
		return err
	}
	fmt.Printf("Created Job %q.\n", result.GetObjectMeta().GetName())
	return nil
}

func (kube *KubeHandler) GetNumOfPods() (num int) {
	pods, err := kube.client.List(metav1.ListOptions{})
	if err != nil {
		return 0
	}
	return len(pods.Items)
}

func (kube *KubeHandler) Delete(jobname string) error {

	deletePolicy := metav1.DeletePropagationForeground
	if err := kube.client.Delete(jobname, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		return err
	}
	fmt.Println("Deleted deployment: " + jobname)
	return nil
}

func (kube *KubeHandler) DeleteCollection(batchid string) error {
	labelPod := labels.SelectorFromSet(map[string]string{"batchid": batchid})
	listPodOptions := metav1.ListOptions{
		LabelSelector: labelPod.String(),
	}
	err := kube.client.DeleteCollection(&metav1.DeleteOptions{}, listPodOptions)

	return err
}

func (kube *KubeHandler) ValidateYaml(yaml string) (Job *argov1alpha.Workflow, err error) {

	job := &argov1alpha.Workflow{}
	dec := k8sYaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(yaml)), 1000)

	if err := dec.Decode(&job); err != nil {

		return nil, err
	}

	return job, nil
}
