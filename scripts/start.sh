DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

set -x

istioctl install --set profile=demo -y
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.8/samples/addons/jaeger.yaml
kubectl label namespace default istio-injection=enabled
kubectl apply -f $DIR/jaeger-go-example.yaml
sleep 10
kubectl get services
kubectl get pods


read  -p "Press any key when all pods are up..."

kubectl apply -f $DIR/jaeger-go-example-gateway.yaml
istioctl analyze
kubectl get svc istio-ingressgateway -n istio-system

export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].hostname}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].port}')

export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT
echo "Run the following command to confirm the example services are up:"
printf "curl -w '\\\n' http://$GATEWAY_URL/ping\n"

read  -p "Press any key if you see 'service-a -> service-b' in STDOUT..."

# Deploy and run jaeger

# First generate some traces
set +x
for i in $(seq 1 100); do curl -s -o /dev/null "http://localhost:8081/ping"; done
set -x
istioctl dashboard jaeger &
