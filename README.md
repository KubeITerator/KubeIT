### KubeIT - A simple iterator framework to parallelize biological analysis in kubernetes

KubeIT is designed to be a framework for the iterative distribution of bioinformatic analyses. 
It exposes a small dedicated API, that is used to simplify workflow scheduling and data handling.
KubeIT uses the cloud native workflow management engine [Argo workflows](https://github.com/argoproj/argo-workflows).


## Contents
- [Dependencies](#dependencies)
- [Installation](#installation)
- [FAQ](#faq)


## Dependencies

- Kubernetes (v)
- S3 Access Credentials
- kubectl ()

## Installation

0. Create a namespace. This namespace must be specified in all following kubectl command with `-n`.
```
kubectl create namespace kubeit
```

1. Create a secret called `s3secret` containing your S3 Credentials (replace `YOUR-S3-KEY` & `YOUR-S3-SECRET`).

```
kubectl create secret generic s3secret --from-literal='AWS_ACCESS_KEY_ID=YOUR-S3-KEY' --from-literal='AWS_SECRET_ACCESS_KEY=YOUR-S3-SECRET' -n kubeit
```

2. Create a secret with KubeITs access token: This token is used for authorization (replace `YOUR-TOKEN`). You can create a new Token on Linux systems with the [openssl](https://www.openssl.org/) command: `openssl rand -base64 20`

```
kubectl create secret generic kubeit-token --from-literal='TOKEN=YOUR-TOKEN' -n kubeit
```

3. Pre-configure KubeIT and Argo configmaps. Download the following file [configmaps.yaml]() . Edit both configmap entries to specify your S3 endpoint, S3 Region and the desired S3 Bucket for data storage. Apply the configmap via:

```
kubectl apply -f confimaps.yaml -n kubeit
```

4. Install Argo. For more information see: [Argo quickstart](https://github.com/argoproj/argo-workflows/blob/master/docs/quick-start.md)

```
kubectl apply -n kubeit -f https://raw.githubusercontent.com/argoproj/argo-workflows/stable/manifests/install.yaml
```

5. Install KubeIT and the default WorkflowTemplates. The first contains the KubeIT deployment and the KubeIT service account and service. The second

```
kubectl apply -n kubeit -f https://raw.githubusercontent.com/argoproj/argo-workflows/stable/manifests/install.yaml
kubectl apply -n kubeit -f https://raw.githubusercontent.com/argoproj/argo-workflows/stable/manifests/install.yaml
```

6. Configure ingress for external access: This last step highly depends on your individual network installation (see [here](https://kubernetes.io/docs/concepts/services-networking/ingress/) for mor information). 
Step 5. creates a service with the name `kubeit-service`. Use this service for configuring external access.
   
## Usage

The KubeIT backend is designed to be used via the [KubeIT CLI](). All backend API routes and their functions / interfaces are explained [here](/tree/dev/API/router/APIDocumentation.md)

### Creating additional WorkflowTemplates:




