---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sr-apps-nse-{{ .Values.nsm.serviceName }}
  namespace: {{ .Release.Namespace }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      networkservicemesh.io/app: "sr-apps-nse-{{ .Values.nsm.serviceName }}"
      networkservicemesh.io/impl: {{ .Values.nsm.serviceName | quote }}
  template:
    metadata:
      labels:
        networkservicemesh.io/app: "sr-apps-nse-{{ .Values.nsm.serviceName }}"
        networkservicemesh.io/impl: {{ .Values.nsm.serviceName | quote }}
    spec:
      containers:
        - name: sr-apps-nse
          image: {{ .Values.registry }}/{{ .Values.org }}/sr-apps_ucnf-nse:{{ .Values.tag }}
          imagePullPolicy: {{ .Values.pullPolicy }}
          env:
            - name: ENDPOINT_NETWORK_SERVICE
              value: {{ .Values.nsm.serviceName | quote }}
            - name: ENDPOINT_LABELS
              value: "app=sr-apps-nse-{{ .Values.nsm.serviceName }}"
            - name: TRACER_ENABLED
              value: "true"
            - name: JAEGER_SERVICE_HOST
              value: jaeger.nsm-system
            - name: JAEGER_SERVICE_PORT_JAEGER
              value: "6831"
            - name: JAEGER_AGENT_HOST
              value: jaeger.nsm-system
            - name: JAEGER_AGENT_PORT
              value: "6831"
            - name: NSREGISTRY_ADDR
              value: "nsmgr.nsm-system"
            - name: NSREGISTRY_PORT
              value: "5000"
            - name: NSE_POD_IP
              valueFrom:
                fieldRef:
                  fieldPath: status.podIP
{{- if .Values.ipam.uniqueOctet }}
            - name: NSE_IPAM_UNIQUE_OCTET
              value: {{ .Values.ipam.uniqueOctet | quote }}
{{- end }}
          securityContext:
            capabilities:
              add:
                - NET_ADMIN
            privileged: true
          resources:
            limits:
              networkservicemesh.io/socket: 1
          volumeMounts:
            - mountPath: /etc/universal-cnf/config.yaml
              subPath: config.yaml
              name: universal-cnf-config-volume
      volumes:
        - name: universal-cnf-config-volume
          configMap:
            name: ucnf-sr-apps-{{ .Values.nsm.serviceName }}
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: ucnf-sr-apps-{{ .Values.nsm.serviceName }}
data:
  config.yaml: |
    endpoints:
    - name: {{ .Values.nsm.serviceName | quote }}
      labels:
        app: "sr-apps-nse-{{ .Values.nsm.serviceName }}"
      vl3:
       ipam:
          defaultPrefixPool: {{ .Values.ipam.prefixPool | quote }}
          prefixLength: 2
          routes: []
       ifName: "endpoint0"
