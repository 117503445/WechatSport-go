
$('#datepicker').val(moment().format('YYYY-MM-DD'))

$.get('/api/people', function (response) {
    for (var i = 0; i < response.length; i++) {
        var option = ['<option>', response[i], '</option>'].join('')
        $('#selectName').append(option)
    }
})

// 基于准备好的dom，初始化echarts实例
var myChart = echarts.init(document.getElementById('main_chart'))
window.onresize = function () {
    myChart.resize()
}

function searchName(name) {
    data = []


    var date = $('#datepicker').val()
    var beginTimeStamp = moment(date).unix()
    var secondsOfDay = 24 * 60 * 60
    var endTimeStamp = beginTimeStamp + secondsOfDay - 1

    var uri = ["/api/record?name=", name, '&beginTimeStamp=', beginTimeStamp, '&endTimeStamp=', endTimeStamp].join('')
    console.log(uri)

    $.get(uri, function (response) {
        if (response == null) {

        } else {
            for (i = 0; i < response.length; i++) {
                var timeStamp = response[i].TimeStamp
                var step = response[i].Step
                var date = new Date(timeStamp * 1000)
                data.push({ value: [date, step] })
            }
        }

        setOption(name)
    });
}

var data = []



function setOption(name) {
    var option = {
        tooltip: {
            trigger: 'axis',
            axisPointer: {
                animation: true
            }
        },
        xAxis: {
            type: 'time',
            splitLine: {
                show: false
            }
        },
        yAxis: {
            type: 'value',
            boundaryGap: [0, '100%'],
            splitLine: {
                show: false
            }
        },
        series: [{
            name: name,
            type: 'line',
            showSymbol: false,
            hoverAnimation: false,
            data: data
        }]
    }
    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option)
}

function search() {
    var name = $('#selectName option:selected').text()
    searchName(name)
}