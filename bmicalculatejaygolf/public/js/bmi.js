$(function () {
    $('#calculate').click(CalculateBmi)
})
function CalculateBmi(){
    
    var height = $('#height').val()
    var weight = $('#weight').val()
    var url = "http://localhost:3000/bmi/calculate?height="+height+"&weight="+weight
    $.getJSON(url, function (responseData) {
       $('#bmi').text(responseData.bmi)
       $('#status').text(responseData.status)
    })
}
function GetApi() {
    var host = "http://localhost:3000/duration"
    var parameter = `?startDay=${$('#startDay').val()}` +
        `&startMonth=${$('#startMonth').val()}` +
        `&startYear=${$('#startYear').val()}` +
        `&endDay=${$('#endDay').val()}` + 
        `&endMonth=${$('#endMonth').val()}` + 
        `&endYear=${$('#endYear').val()}`
    var url = host + parameter
    $.getJSON(url, function (responseData) {
        $('#resultDay').html(responseData.days)
    })
}