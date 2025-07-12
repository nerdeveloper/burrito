# Build the Burrito controller image locally
docker_build('burrito-controller', '.', dockerfile='Dockerfile')

# Deploy Burrito using Helm chart
k8s_yaml(helm(
    './deploy/charts/burrito',
    name='burrito',
    namespace='burrito-system',
    values=['./deploy/charts/burrito/values-dev.yaml'],
    set=[
        'global.deployment.image.repository=burrito-controller',
        'global.deployment.image.pullPolicy=Never',  # Use local image
        'config.burrito.runner.image.repository=burrito-controller',
        'config.burrito.runner.image.tag=tilt-98fc8ad1e6d63da6',  # Use same tag as controller
        'config.burrito.runner.image.pullPolicy=Never',  # Use local image for runners too
        'config.burrito.datastore.storage.mock=true',  # Use in-memory storage for testing
    ]
))

# Create the burrito-system namespace if it doesn't exist
k8s_yaml(blob('''
apiVersion: v1
kind: Namespace
metadata:
  name: burrito-system
'''))

# Create CRDs first (exclude kustomization.yaml)
local_resource(
    'apply-crds',
    'kubectl apply -f manifests/crds/config.terraform.padok.cloud_terraformlayers.yaml -f manifests/crds/config.terraform.padok.cloud_terraformpullrequests.yaml -f manifests/crds/config.terraform.padok.cloud_terraformrepositories.yaml -f manifests/crds/config.terraform.padok.cloud_terraformruns.yaml',
    deps=['manifests/crds/'],
    labels=['setup']
)

# Watch for Go code changes in controller
k8s_resource('burrito-controllers', 
    port_forwards=['8091:8081'],  # Health endpoint on different port
    labels=['burrito']
)

k8s_resource('burrito-datastore', 
    port_forwards=['8090:8080'],  # Main datastore port on different port
    labels=['burrito']
)

k8s_resource('burrito-server', 
    port_forwards=['8092:8080'],  # Server on different local port
    labels=['burrito']
)


print("Tilt setup complete!")
print("1. Run 'tilt up' to start the dev environment")
print("2. Generate a GitHub token with: gh auth refresh -h github.com -s repo")
print("3. Create GitHub secret: kubectl create secret generic github-creds --from-literal=githubToken=YOUR_TOKEN -n burrito-system")
print("4. Apply test CRDs to start testing")
