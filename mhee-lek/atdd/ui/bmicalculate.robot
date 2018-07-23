*** Setting ***
Library    SeleniumLibrary
Suite Setup    Open Browser    ${URL}    ${BROWSER}
Test Setup    Go To    ${URL}
Suite Teardown    Close Browser

*** Variable ***
${URL}    http://localhost:4321/bmi/
${BROWSER}    chrome

*** Testcases ***
คำนวณค่า BMI ความสูง 170 หนัก 80
    Input Text    height    170
    Input Text    weight    80
    Click Button    calculate
    Page Should Contain Element    bmi    27.7
    Page Should Contain Element    status    อ้วน