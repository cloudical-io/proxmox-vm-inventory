package web

var table string = `
<!DOCTYPE html>
<html>
    <head>
        <style>
            table {
                font-family: arial, sans-serif;
                border-collapse: collapse;
                width: 100%;
            }

            td, th {
                border: 1px solid #dddddd;
                text-align: left;
                padding: 8px;
            }

            tr:nth-child(even) {
                background-color: #dddddd;
            }

            td:nth-child(4), th:nth-child(4){
                min-width: auto;
            }

            td, th{
                min-width: max-content;
            }
        </style>
    </head>
    <body>
        {{ range $key, $value := . }}
        <h3>Cluster: {{ $key }}</h3>
        <table>
            <tr>
                <th>Name</th>
                <th>Vmid</th>
                <th>Status</th>
                <th>Networks</th>
                <th>Tags</th>
                <th>Cpu</th>
                <th>Maxmem</th>
                <th>Maxdisk</th>
                <th>Lock</th>
                <th>Pid</th>
                <th>Qmpstatus</th>
                <th>Runningmachine</th>
                <th>Runningqemu</th>
                <th>Uptime</th>
            </tr>
            {{ range $element := $value }}
            <tr>
                <td>{{ .Name }}</td>
                <td>{{ .Vmid }}</td>
                <td>{{ .Status }}</td>
                <td>{{ joinNetworks .Networks }}</td>
                <td>{{ .Tags }}</td>
                <td>{{ .Cpu }}</td>
                <td>{{ .Maxmem }}</td>
                <td>{{ .Maxdisk }}</td>
                <td>{{ .Lock }}</td>
                <td>{{ .Pid }}</td>
                <td>{{ .Qmpstatus }}</td>
                <td>{{ .Runningmachine }}</td>
                <td>{{ .Runningqemu }}</td>
                <td>{{ .Uptime }}</td>
            </tr>
            {{ end }}
        </table>
        {{ end }}
    </body>
</html>
`
