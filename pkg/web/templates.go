package web

var table string = `
<!DOCTYPE html>
<html>
    <body>
    </body>
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
	</style>
    <body>
	<table>
	{{range $key, $value := .}}
    <h3>{{$key}}</h3>
    <table>
        <tr>
            <th>Name</th>
            <th>Status</th>
            <th>Vmid</th>
            <th>Cpu</th>
            <th>Lock</th>
            <th>Maxdisk</th>
            <th>Maxmem</th>
            <th>Pid</th>
            <th>Qmpstatus</th>
            <th>Runningmachine</th>
            <th>Runningqemu</th>
            <th>Tags</th>
            <th>Uptime</th>
            <th>Networks</th>
        </tr>
        {{range $element := $value}}
        <tr>
            <td>{{.Name}}</td>
            <td>{{.Status}}</td>
            <td>{{.Vmid}}</td>
            <td>{{.Cpu}}</td>
            <td>{{.Lock}}</td>
            <td>{{.Maxdisk}}</td>
            <td>{{.Maxmem}}</td>
            <td>{{.Pid}}</td>
            <td>{{.Qmpstatus}}</td>
            <td>{{.Runningmachine}}</td>
            <td>{{.Runningqemu}}</td>
            <td>{{.Tags}}</td>
            <td>{{.Uptime}}</td>
            <td>{{joinNetworks .Networks}}</td>
        </tr>
        {{end}}
    </table>
    {{end}}
</html>
`
