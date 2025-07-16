# Runal with OSC

This was a quick test to show how [Runal](https://empr.cl/runal/) 
could be controlled by an OSC client on a phone. 

## Slider Example

A circle's size changes based on a slider on the phone.
I used the [Osc Controller](https://play.google.com/store/apps/details?id=com.ffsmultimedia.osccontroller&hl=en-US)
Android app, which has `/slider1` as the default message for the first slider.

## Multitouch Example

I used the "Sensors2OSC" Android app, which can send multitouch
data. You can also use it to send data from all the other sensors
on the phone. The app is available on [F-Droid](https://f-droid.org/packages/org.sensors2.osc/)
or as an [apk on Codeberg](https://codeberg.org/Residuum/Sensors2OSC/releases).

The app doesn't show you what messages are being sent, but my [oscdump](https://github.com/rahji/oscdump)
program shows that it is sending messages like this for (fingers 1, 2, and 3):

```osc
/touch1 ,ff 0.23026891 0.4360344
/touch2 ,ff 0.4832411 0.6755878
/touch3 ,ff 0.7616663 0.5655826
  
```

## Windows Firewall

I don't remember having to change the Windows Firewall in the past (as in last week), but
yesterday I couldn't receive UDP packets on any port in Windows. I had
to open *Windows Defender Firewall with Advanced Security* and open UDP
port 9999. Now it works in both Windows and WSL2, as long as I use that port in my
OSC client app. 🤷🏽‍♂️

