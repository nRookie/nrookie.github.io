``` shell
my-mongo-pvc.yaml

# my-mongo-pvc.yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-mongo-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
```





``` shell
# my-mongo-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-mongo
spec:
  selector:
    matchLabels:
      app: my-mongo
      environment: my
  replicas: 1
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: my-mongo
        environment: my
    spec:
      hostname: my-mongo
      volumes:
        - name: my-mongo-persistent-storage
          persistentVolumeClaim:
            claimName: my-mongo-pvc
      containers:
        - name: my-mongo
          image: mongo:3.6
          imagePullPolicy: Always
          ports:
            - containerPort: 27017
          volumeMounts:
            - name: my-mongo-persistent-storage
              mountPath: /data/db
```



``` shell
# my-mongo-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: my-mongo
spec:
  ports:
  - port: 27017
  selector:
    app: my-mongo
  clusterIP: None
```



``` shell
kubectl apply -f ./my-mongo-pvc.yaml \
  -f ./my-mongo-deployment.yaml \
  -f ./my-mongo-service.yaml
```





# 2 | Test connection



