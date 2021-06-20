# Sample-controller
It's the controller to create resources thar outputs message you set on manifest with 'Hello'.

## deploy
Before deploy the resources, you need to execute the commands below.

1. Install CRD.

```
$ make install
```

2. Run custom controller on your terminal.

```
$ make run
```

Then you can deploy the resources with following manifest.

```yaml
apiVersion: samplecontroller.my.domain/v1
kind: Bar
metadata:
  name: bar-sample-1
spec:
  # Add fields here
  message: Bar
```

After that you can see outputs with `kubectl get`.

```
% kubectl get ba
NAME           MESSAGE
bar-sample-1   Hello Bar
```
