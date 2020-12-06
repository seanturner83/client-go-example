# client-go-example

A really simple example of using client-go for accessing the Kubernetes API, building a container with the binary in, and spinning it up in a namespace with the relevant RBAC access.

## Installation

If you're happy to use my image, you can `kubectl apply -f k8s/`
Of course you may wish to edit the yaml files therein and have them operate on some other namespace - this is just an example after all.

If you want to build your own image, there's a Dockerfile in the root - this should do everything for you, just build, tag, and push.

## Contributing

Seriously? You could add some tests, kustomize the namespace bit, add Helm, or maybe just make it a bit more useful.

## License

`do what you like`
