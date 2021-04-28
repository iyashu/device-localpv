{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "devicelocalpv.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified localpv provisioner name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "devicelocalpv.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "devicelocalpv.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}


{{/*
Create the name of the service account for controller
*/}}
{{- define "devicelocalpv.deviceController.serviceAccountName" -}}
{{- if .Values.serviceAccount.deviceController.create }}
{{- default (include "devicelocalpv.fullname" .) .Values.serviceAccount.deviceController.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.deviceController.name }}
{{- end -}}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "devicelocalpv.deviceNode.serviceAccountName" -}}
{{- if .Values.serviceAccount.deviceNode.create }}
{{- default (include "devicelocalpv.fullname" .) .Values.serviceAccount.deviceNode.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.deviceNode.name }}
{{- end -}}
{{- end -}}

{{/*
Define meta labels for openebs device-localpv components
*/}}
{{- define "devicelocalpv.common.metaLabels" -}}
chart: {{ template "devicelocalpv.chart" . }}
heritage: {{ .Release.Service }}
openebs.io/version: {{ .Values.release.version | quote }}
role: {{ .Values.role | quote }}
{{- end -}}

{{/*
Create match labels for openebs device-localpv controller
*/}}
{{- define "devicelocalpv.deviceController.matchLabels" -}}
app: {{ .Values.deviceController.componentName | quote }}
release: {{ .Release.Name }}
component: {{ .Values.deviceController.componentName | quote }}
{{- end -}}

{{/*
Create component labels for devicelocalpv controller
*/}}
{{- define "devicelocalpv.deviceController.componentLabels" -}}
openebs.io/component-name: {{ .Values.deviceController.componentName | quote }}
{{- end -}}


{{/*
Create labels for openebs device-localpv controller
*/}}
{{- define "devicelocalpv.deviceController.labels" -}}
{{ include "devicelocalpv.common.metaLabels" . }}
{{ include "devicelocalpv.deviceController.matchLabels" . }}
{{ include "devicelocalpv.deviceController.componentLabels" . }}
{{- end -}}

{{/*
Create match labels for openebs device-localpv node daemon
*/}}
{{- define "devicelocalpv.deviceNode.matchLabels" -}}
name: {{ .Values.deviceNode.componentName | quote }}
release: {{ .Release.Name }}
{{- end -}}

{{/*
Create component labels openebs device-localpv node daemon
*/}}
{{- define "devicelocalpv.deviceNode.componentLabels" -}}
openebs.io/component-name: {{ .Values.deviceNode.componentName | quote }}
{{- end -}}


{{/*
Create labels for openebs device-localpv node daemon
*/}}
{{- define "devicelocalpv.deviceNode.labels" -}}
{{ include "devicelocalpv.common.metaLabels" . }}
{{ include "devicelocalpv.deviceNode.matchLabels" . }}
{{ include "devicelocalpv.deviceNode.componentLabels" . }}
{{- end -}}
