apiVersion: v1
kind: Secrettest
metadata:
  name: azure-blobstore-secret
  namespace: kyma-installer
type: Opaque
data:
  shared_key: "__KYMA_RELEASES_AZURE_BLOBSTORE_KEY__"
