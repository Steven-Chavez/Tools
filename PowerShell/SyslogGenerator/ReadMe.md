# UDP-SyslogGenerator.ps1

## Overview
`UDP-SyslogGenerator.ps1` is a PowerShell script designed to send a series of syslog messages to a specified syslog server using UDP. This script is particularly useful for testing and validating the reception of syslog messages by a syslog server or Security Information and Event Management (SIEM) system.

## Use Case
This script is ideal for security professionals and system administrators who need to:
- Test the configuration and performance of syslog servers.
- Validate the integration of syslog messages with SIEM systems.
- Simulate syslog traffic for troubleshooting and monitoring purposes.

## Parameters
- **SyslogServer**: Specifies the address or hostname of the syslog server to which the syslog messages will be sent.
- **Port**: Specifies the port number to use when sending syslog messages over UDP. The default port is 514.
- **NumOfEvents**: Specifies the number of syslog events to send. The default is 20.
- **Identifier**: Specifies an optional identifier to include in the syslog messages. This is useful for searching for the test logs.

## Example
```powershell
.\UDP-SyslogGenerator.ps1 -SyslogServer "10.0.0.1" -Port 514 -NumOfEvents 20 -Identifier "Test"
