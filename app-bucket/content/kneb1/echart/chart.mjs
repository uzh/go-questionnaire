// https://github.com/apache/echarts
// https://github.com/ecomfe/echarts-stat
// https://github.com/ecomfe/awesome-echarts
// https://echarts.apache.org/examples/en/editor.html?c=bar-histogram

// import simply does not work; possibly because of stackoverflow.com/questions/71022803
// import * as echarts from './echarts.min.js';
// import ecStat from './ecStat.js';



var chartDom = document.getElementById('chart_container');
// console.log(chartDom);
var myChart = echarts.init(chartDom);
var option;



var colorPalette = ['#d87c7c', '#919e8b', '#d7ab82', '#6e7074', '#61a0a8', '#efa18d', '#787464', '#cc7e63', '#724e58', '#4b565b'];
function getColor() {
    let idx = colorPalette.length % counterDraws;
    return colorPalette[idx];
}


var ds1Example = {
    source: [
        [10.3, 143],
        [10.6, 214],
        [10.8, 251],
        [10.0, 176],
        [10.1, 221],
        [10.2, 188],
        [10.4, 191],
        [10.0, 196],
        [10.9, 177],
        [10.9, 153],
        [10.3, 201],
    ]
};

var ds2Example = {
    source: [
        [ 25, 0],
        [ 75, 0],
        [125, 0],
        [175, 0],
        [225, 0],
        [275, 0],
    ]
};


var ds1 = {
    source: []
};
var counterDrawsInit = 4 ;
var counterDraws = counterDrawsInit;  // counter for getData




var mn = 0.0; // mean
var sd = 1.0; // standard deviation

var mn = 153.0; // mean
var sd = 15.89; // standard deviation

var normDist = new jsstats.NormalDistribution(mn, sd);


// histogram config
// ======================

const  w = 20;   // width
const wh = 10;   // width half

var maxXHisto = 0;

function getMax() {
    return maxXHisto+2; 
}



var ds2 = {
    source: []
};
for (let i = 0; i <= 300/w; i++) {
    ds2.source.push([wh + i*w, 0]);    
}

// console.log(ds2.source);

function getData() {

    counterDraws++;

    for (let i = ds1.source.length; i < (counterDraws+1); i++) {
        let linDraw = Math.random();
        let draw  = normDist.invCumulativeProbability(linDraw)
        // console.log(`   lin draw ${linDraw} => draw  ${draw}`);

        let subAr = ["draw", draw];
        ds1.source.push(subAr);
    }

    // console.log(`counterDraws ${counterDraws} - ds1a: `, ds1a.source );


    // histogram data


    let i0 = ds1.source.length - 1
    if (counterDraws == counterDrawsInit+1) {
        i0 = 0;
    }
    for (let i = i0; i < ds1.source.length; i++) {
        let val    = Math.floor(ds1.source[i][1]);
        let binId  = Math.round(val/w)*w + wh;
        let binIdx = (binId - wh) / w;
        // console.log(`   val ${val} => binId ${binId} - => binIdx ${binIdx}`);
        ds2.source[binIdx][1]++;
        if (ds2.source[binIdx][1] > maxXHisto) {
            maxXHisto = ds2.source[binIdx][1];
        }

    }

    // console.log(`counterDraws ${counterDraws} - ds2a: `, ds2a.source);

    return [
        ds1,
        ds2,
        {
            transform: {
                type: 'ecStat:histogram',
                // print: true,
                config: { dimensions: [1] }
            }
        },
    ];
}

option = {
    dataset: getData(),
    tooltip: {},
    grid: [
        {
            top:   '04%',
            right: '60%',
        },
        {
            top:    '04%',
            left:   '56%',
        }
    ],
    xAxis: [
        {
            type:  'value',
            type:  'category',
            // do not include zero position 
            scale:  true,  
            gridIndex: 0
        },
        {
            scale: true,
            gridIndex: 1,
            inverse: true,

            min: 0,
            max: function(){
                return getMax();
            },

        }
    ],
    yAxis: [
        {
            gridIndex: 0,
            min: 0,
            max: 300,
        },
        {
            gridIndex: 1,
            // min: 0,
            // max: 350,
            // must be category: https://github.com/apache/echarts/issues/15960
            //      or https://echarts.apache.org/en/option.html#series-custom
            type: 'category',
            // axisTick:  { show: false },
            // axisLabel: { show: false },
            // axisLine:  { show: false },

            axisLine: {
                // necessary for position: right to take effect
                onZero: false,
            },
            position: 'right',
        }
    ],
    series: [
        {
            name: 'random draws',
            type: 'scatter',
            xAxisIndex: 0,
            yAxisIndex: 0,
            encode: { tooltip: [0, 1] },
            symbol: 'emptyCircle',
            symbol: 'circle',
            // symbolOffset only works for the entire series
            //   symbolOffset: [  -33, 10],
            //   symbolOffset: [ Math.floor((Math.random() *  44)) -22],
            symbolSize: function (value, params) {
                // console.log(`symbolSize`, params.data);
                // console.log(`symbolSize`, params.dataIndex, counterDraws);
                let a1 = params.dataIndex + 1;
                let a2 = counterDraws;
                if (a1 == a2) {
                    return 7;
                }
                params.color = '#919e8b';  // does not affect
                return 3;
                // return value;
            },
            itemStyle: {
                // borderWidth: 3,
                // borderColor: '#EE6666',

                // color function not possible
                // color: function (value, params) {
                //     return '#919e8b';
                // },
                // color: 'yellow',


                opacity: 0.3,
            },

            // color does not work as symbolSize 
            // color: function (value, params) {
            //     return getColor();
            // },
            // color: getColor(),

            datasetIndex: 0,
        },
        {
            name: 'histogram',
            type: 'bar',
            xAxisIndex: 1,
            yAxisIndex: 1,
            barWidth: '99.3%',
            barWidth: '2px',
            label: {
                show: true,
                position: 'right',
                position: 'center',

            },
            encode: { x: 1, y: 0, itemName: 4 },
            datasetIndex: 1
        }
    ],
};

option && myChart.setOption(option);


setInterval(() => {
    myChart.setOption({
        dataset: getData(),
        // series: {
        //     data: makeRandomData()
        // }
    });
}, 200);