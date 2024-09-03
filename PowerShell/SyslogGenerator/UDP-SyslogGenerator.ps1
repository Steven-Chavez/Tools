<#
.SYNOPSIS
    This script sends a series of syslog messages to a specified syslog server using UDP.

.DESCRIPTION
    This PowerShell script sends syslog messages to a syslog server over UDP. You can specify the syslog server, port, number of events, and an optional identifier for the syslog messages.

.PARAMETER SyslogServer
    Specifies the address or hostname of the syslog server to which the syslog messages will be sent.

.PARAMETER Port
    Specifies the port number to use when sending syslog messages over UDP. The default port is 514.

.PARAMETER NumOfEvents
    Specifies the number of syslog events to send. The default is 20.

.PARAMETER Identifier
    Specifies an optional identifier to include in the syslog messages. This can be used to label or categorize the syslog messages.

.EXAMPLE
    .\UDP-SyslogGenerator.ps1 -SyslogServer "10.0.0.1" -Port 514 -NumOfEvents 20 -Identifier "Test"

    This example sends 20 syslog events to the syslog server at '10.0.0.1' on port 514 with the identifier "Test".

.NOTES
    Author: Steven Chavez
    Date: 10/5/2023
    Version: 1.0
    Filename: UDP-SyslogGenerator.ps1
#>

param(
    [string] $SyslogServer,       # The address or hostname of the syslog server.
    [int] $Port = 514,            # The port number for UDP communication (default is 514).
    [int] $NumOfEvents = 20,      # The number of syslog events to send (default is 20).
    [string] $Identifier = "Test-001"     # An optional identifier to include in syslog messages.
)

for ($i = 0; $i -lt $NumOfEvents; $i++) {
    # Construct the syslog message with the event number and optional identifier.
    $SyslogMessage = "<14>Test syslog message from PowerShell - Event " + $i + " " + $Identifier
    
    # Create a UDP client for sending syslog messages.
    $UDPClient = New-Object System.Net.Sockets.UdpClient
    
    # Convert the syslog message to bytes.
    $SyslogBytes = [System.Text.Encoding]::ASCII.GetBytes($SyslogMessage)
    
    # Send the syslog message to the specified syslog server and port.
    $UDPClient.Send($SyslogBytes, $SyslogBytes.Length, $SyslogServer, $Port)
    
    # Close the UDP client.
    $UDPClient.Close()
    
    # Pause for 1 second before sending the next syslog message.
    Start-Sleep -Seconds 1
}
