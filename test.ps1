$port = 8080  # Port sur lequel écouter
$server = New-Object System.Net.Sockets.TcpListener("127.0.0.1", $port) # Écoute sur localhost
$server.Start()


$process = Start-Process -FilePath ".\winpackgo.exe" -NoNewWindow -PassThru
$client = $server.AcceptTcpClient()
$stream = $client.GetStream()
$reader = New-Object System.IO.StreamReader($stream)
$response = $reader.ReadLine()
Write-Host "Message reçu de Go : $response"

$reader.Close()
$client.Close()
$server.Stop()

$process.WaitForExit()