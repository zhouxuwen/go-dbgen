
{{.TableComment}}
type {{.Table}} struct {
{{range $j, $item := .Fields}}{{$item.Name}}       {{$item.Type}}    {{$item.FormatFields}}        {{$item.Remark}}
{{end}}
}