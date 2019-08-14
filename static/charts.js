


// 基于准备好的dom，初始化echarts实例
var myChart = echarts.init(document.getElementById('main'));
window.onresize = function () {
    myChart.resize()
}
searchName("117503445")

function searchName(name) {
    data = []
    $.get("http://127.0.0.1/api/record?name=" + name, function (response) {
        //console.log(response)
        for (i = 0; i < response.length; i++) {
            var timeStamp = response[i].TimeStamp
            var step = response[i].Step
            var date = new Date(timeStamp * 1000);
            data.push({ value: [date, step] });
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
    };
    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
}

function search() {
    var name = $('#selectName option:selected').text()
    searchName(name)
    console.log(name)
}