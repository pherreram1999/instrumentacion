const int pinInput = A0;
float voltageAnalog = 0; 
float voltage = 0;

void setup() {
  pinMode(pinInput,INPUT);
  Serial.begin(9600);

}

void loop() {
  // put your main code here, to run repeatedly:
  voltageAnalog = analogRead(pinInput);
  voltage = (voltageAnalog * 5.0) /1023.0;
  delay(500);
  Serial.println(voltage);
}
