{{- define "neo4j.nameOverride" -}}
  {{- .Values.neo4j.name | default .Release.Name }}
{{- end -}}
