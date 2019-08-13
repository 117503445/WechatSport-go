// 基于准备好的dom，初始化echarts实例
var myChart = echarts.init(document.getElementById('main'));

var tdata = []

$.get("http://127.0.0.1/api/record?name=冯凯迪", function (response) {
    //console.log(response)
    for (i = 0; i < response.length; i++) {
        var timeStamp = response[i].TimeStamp
        var step = response[i].Step
        var date = new Date(timeStamp * 1000);
        console.log(date, step)
        tdata.push({ value: [date, step] });
    }
    console.log(data)
    console.log(tdata)
    setOption()
});

function setOption() {
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
            name: '模拟数据',
            type: 'line',
            showSymbol: false,
            hoverAnimation: false,
            data: tdata
        }]
    };
    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
}


function randomData() {
    now = new Date(+now + oneDay);
    value = value + Math.random() * 21 - 10;
    return {
        value: [
            [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('/'),
            100
        ]
    }
}

var data = [];
var now = +new Date(2019, 8, 13);
var oneDay = 24 * 3600 * 1000;
var value = Math.random() * 1000;
for (var i = 0; i < 10; i++) {
    data.push(randomData());
}
