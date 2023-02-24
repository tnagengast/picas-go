# Picas-Go

An fun lil API that turns text into ASCII art, like Picasso! 🎨

## Notes

Set CLI env variables
```sh
export REPLICATED_APP=picas-go
export REPLICATED_API_TOKEN=***
```

Connect to the Replicated private registry
```sh
docker login registry.replicated.com
```

Attach to minikube docker env
```sh
eval $(minikube docker-env)
```

Dettach to minikube docker env
```sh
eval $(minikube docker-env --unset)
```

Build the image
```sh
docker build -t registry.replicated.com/picas-go/api:1.0.0 .
docker build --platform=linux/amd64 -t registry.replicated.com/picas-go/api:1.0.2-amd64 .
```

Run the container and remove when stopped
```sh
docker run -it --rm -p 8080:8080 registry.replicated.com/picas-go/api:1.0.0
```

List docker images in current env
```sh
docker images
```

Push the image the the remote registry
```sh
docker push registry.replicated.com/picas-go/api:1.0.0
```

Lint releases
```sh
replicated release lint --yaml-dir=manifests
```

Cut a new release
```sh
replicated release create \
  --version 1.0.2 \
  --release-notes "Build image for linux" \
  --promote Unstable \
  --yaml-dir manifests
```

List releases
```sh
replicated release ls
```

Show channel information
```sh
replicated channel inspect Unstable
```
