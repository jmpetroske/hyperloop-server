# Communication between server and web end
All web -> server data should be submitted using POST requests

## Endpoints
- `/mission`: submit mission variables. See [Setting mission parameters](#setting-mission-parameters)
- `/arm`: arm the pod
- `/start`: make the pod go
- `/command`: For use in god mode. see [god mode commands](#god-mode-commands)
- `/abort`: Abort the mission
- `/dataWebSocket`: Open a websocket for streaming data

## Setting mission parameters
Set 3 variables, all in float format: `distance`, `pressure`, and `topSpeed`.

## God mode commands
Set the `command` variable to a value between 0 and 5.
| Value | Meaning             |
|-------|---------------------|
| 0     | engageBreaks        |
| 1     | disengageBreaks     |
| 2     | engageSolenoids     |
| 3     | disengageSolenoids  |
| 4     | engageBallValves    |
| 5     | disengageBallValves |

## Example websocket data
```javascript
{"timestamp":100000,"mode":3,"pressureThrusterTank":88.82,"pressureBrakingTank":65.83,"pressureBrakingExtend":56.13,"pressureBrakingCompress":76.29,"temperatureThrusterLeftTank":53.54,"temperatureThrusterRightTank":63.85,"temperatureThrusterNozzle":95.88,"temperatureBattery":13.46,"temperatureBrakingTank":51.1,"temperatureBrakingExtend":15.8,"temperatureBrakingCompress":10.73,"speed":59.631397,"distance":70.96314,"vibration1":78.79,"vibration2":63.15,"vibration3":68.73,"vibration4":96.57,"actuationThruster":true,"actuationRelieveValve":true,"actuationSafetyValve":true,"actuationBrakingValve":false,"currentThrusterLeftValve":4.38,"currentThrusterRightValve":27.92,"currentThrusterRelieveValve":20.33,"currentBrakingValve":7.21,"currentNAP":48.2,"currentBattery":88.49,"voltageThrusterLeftValve":40,"voltageThrusterRightValve":75.71,"voltageThrusterRelieveValve":78.43,"voltageBrakingValve":78.41,"voltageNAP":28.81,"voltageBattery":58.13,"orientationx":90.1,"orientationy":87.9,"orientationz":76.49,"accelerationx":34.64,"accelerationy":12.4,"accelerationz":96.29,"imuStatus":1,"brakeStatus":0,"propulsionStatus":1,"emergencyBrake":false}
```
