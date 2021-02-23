# Short guideline for creating new WorkflowTemplates

Examples for workflow templates can be found [here](templates/).

Templates should be modular, they should define all needed parameters in the beginning and then use then at appropiate place.
WorkflowTemplates can have the same structure as regular Argo Workflows.

An short example for a template is:


```yaml
apiVersion: argoproj.io/v1alpha1
kind: WorkflowTemplate
metadata:
  name: init-template
spec:
  arguments:
    parameters:
      - name: INPUTDATA
        value: ""
      - name: COMMAND
        value: ""
      - name: VOLUMENAME
        value: ""
      - name: IMAGE
        value: ""

# Workflow template and order
  templates:
    - name: init
      inputs:
        parameters:
          - name: INPUTDATA
          - name: COMMAND
          - name: VOLUMENAME
          - name: IMAGE
        artifacts:
          - name: inputfile
            path: /data/inputfile
            mode: 0755
            http:
              url: "{{inputs.parameters.INPUTDATA}}"
      container:
        image: "{{inputs.parameters.IMAGE}}"
        securityContext:
          runAsUser: 10001
        resources:
          limits:
            memory: "2Gi"
            cpu: "2000m"
          requests:
            memory: "1Gi"
            cpu: "1000m"
        command: [sh, -c]
        args: ["{{inputs.parameters.COMMAND}}"]
        volumeMounts:
        - name: "{{inputs.parameters.VOLUMENAME}}"
          mountPath: /data/
        


```

The arguments section describes initial config parameters. The input section of a template describes used parameters and artifacts. Input parameters can be referenced via the term `{{inputs.parameters.PARAMETERNAME}}`.