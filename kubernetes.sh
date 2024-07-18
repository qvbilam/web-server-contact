servername=contact
serverPort=9705
targetPort=9705

# 申请配置
kubectl apply -f ${servername}.config.yaml
kubectl apply -f ${servername}.secret.yaml
kubectl apply -f ${servername}.deployment.yaml
kubectl apply -f ${servername}.server.yaml
# 开放端口
kubectl port-forward service/${servername}-web-server ${serverPort}:${targetPort} -n default