*** Settings ***
Library    SeleniumLibrary
Suite Setup    Open Browser    http://localhost:3000/bmi/    chrome
Suite teardown    ปิด Browser
*** Variables ***

*** TestCases ***
คำนวน bmi ของพี่หมี 
    เปิด Browser 
    ใส่ส่วนสูง    170
    ใส่น้ำหนัก    80
    กดปุ่มคำนวน 
    จะได้ค่า bmi และ bmistatus    27.7    อ้วน

คำนวน bmi ของเล็ก
    เปิด Browser 
    ใส่ส่วนสูง    170
    ใส่น้ำหนัก     50
    กดปุ่มคำนวน 
    จะได้ค่า bmi และ bmistatus    17.3    ผอม 

*** Keywords ***
เปิด Browser 
    go to     http://localhost:3000/bmi    
ใส่ส่วนสูง
    [Arguments]    ${height}    
    Input Text    id=height    ${height}
ใส่น้ำหนัก
    [Arguments]    ${weight}    
    Input Text    id=weight    ${weight}
กดปุ่มคำนวน
    Click Button    id=calculate 
จะได้ค่า bmi และ bmistatus 
    [Arguments]    ${bmi}   ${status}
    Wait Until Page Contains     ${bmi}   
    Wait Until Page Contains     ${status}    
ปิด Browser 
    Close Browser