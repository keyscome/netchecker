{{- /* report.tpl */ -}}
Network Connection Report - {{ .Timestamp }}

{{ range $service, $result := .Results }}
Service: {{ $service }}
  Success:
  {{- if $result.Success }}
  {{- range $result.Success }}
    - {{ . }}
  {{- end }}
  {{- else }}
    (无成功连接)
  {{- end }}

  Failures:
  {{- if $result.Failure }}
  {{- range $result.Failure }}
    - {{ . }}
  {{- end }}
  {{- else }}
    (无失败连接)
  {{- end }}

{{ end }}
