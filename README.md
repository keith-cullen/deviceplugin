# DevicePlugin

A skeleton k8s device plugin project.

## Instructions

1. Build a stand-alone executable

        $ go build

2. Run a stand-alone executable

        $ ./deviceplugin

3. Build a Docker image

        $ docker build --tag localhost:5000/deviceplugin:latest .

4. Run a Docker container

        $ docker run -d localhost:5000/deviceplugin:latest

5. Push the Docker image to a local Docker registry

        $ docker push localhost:5000/deviceplugin:latest

6. Create a Kubernetes pod

        $ kubectl apply -f devicepluginpod.yaml

7. List advertised device plugin resources

        $ kubectl describe node <node-name>
        $ kubectl get nodes -o json | jq '.items[].status.allocatable'
