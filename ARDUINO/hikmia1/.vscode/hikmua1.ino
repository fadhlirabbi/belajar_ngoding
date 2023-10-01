
// WIFI
#include <WiFi.h>
String wifiSSID = "HIKMO";
String wifiPassword = "hikmoapake";

// BOT telegram
#include <ArduinoJson.h>
#include <CTBot.h>
CTBot myBot;
TBMessage tMessage;
String token = "6037416728:AAHoXREcDcXkZOY_M63UzuzJfH_kpuXi0RM"; //token API Bot Telegram

// Komponen
const int pirPin = 2;         // Pin sensor PIR (HC-SR501)
const int doorSwitchPin = 4;  // Pin magnetic switch pintu


// Deklarasi Fungsi
void connectWifi();
void loginTelegram();

void setup() {

  Serial.begin(9600); 
  
  // Memanggil fungsi
  connectWifi();
  loginTelegram();


}

void loop()
{
  
  if (myBot.getNewMessage(tMessage))
  {
    Serial.println(tMessage.text);
    if (tMessage.messageType == CTBotMessageText)
    {
      if (tMessage.text.equalsIgnoreCase("/start"))
      {
        String reply = "";
        reply += "Hii, welcome !!! \n";
        reply += "Ini adalah bot keamanan pintu, jika pintu terbuka atau ada gerakan didepan sensor, maka bot akan mengirimkan notifikasi kepada anda \n";
        myBot.sendMessage(tMessage.sender.id, reply);
      }

      else
      {
        String reply = "";
        reply += "perintah tersebut tidak valid, silahkan klik /start untuk melihat petunjuk penggunaan\n";
        myBot.sendMessage(tMessage.sender.id, reply);
      }
    }
  }
  
  else if (digitalRead(pirPin) == HIGH)
  {
    String reply = "";
    reply += "Gerakan Terdeteksi oleh sensor";
    myBot.sendMessage(tMessage.sender.id, reply);
  }

  else if (digitalRead(doorSwitchPin) == LOW)
  {
    String reply = "";
    reply += "pintu tidak tertutup";
    myBot.sendMessage(tMessage.sender.id, reply);
  }
}

void connectWifi()
{
  Serial.println("Connecting To Wifi");
  WiFi.begin(wifiSSID.c_str(), wifiPassword.c_str());

  while (WiFi.status() != WL_CONNECTED)
  {
    Serial.print(".");
    delay(500);
  }
}

void loginTelegram()
{
  Serial.println("logging in...");
  while (!myBot.testConnection())
  {
    myBot.setTelegramToken(token);
    delay(1000);
  }
  Serial.println("Telegram connection OK!");
}