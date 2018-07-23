*** Settings ***
Library    SeleniumLibrary

*** Variables ***

*** TestCases ***
คำนวน bmi ของพี่หมี 
    เปิด Browser 
    ใส่ส่วนสูง  
    ใส่น้ำหนัก
    กดปุ่มคำนวน 
    จะได้ค่า bmi และ bmistatus 
    ปิด Browser

*** Keywords ***
เปิด Browser 
    Open Browser    http://localhost:3000/bmi    chrome
ใส่ส่วนสูง
    Input Text    id=height    170
ใส่น้ำหนัก
    Input Text    id=weight    80
กดปุ่มคำนวน
    Click Button    id=calculate 
จะได้ค่า bmi และ bmistatus 
    Wait Until Page Contains     27.7    
    Wait Until Page Contains     อ้วน    
ปิด Browser 
    Close Browser 