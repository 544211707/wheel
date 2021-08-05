<template>
  <LuckyWheel
    style="margin: 5% auto;"
    ref="LuckyWheel"
    width="300px"
    height="300px"
    :prizes="prizes"
    :default-style="defaultStyle"
    :blocks="blocks"
    :buttons="buttons"
    @start="startCallBack"
    @end="endCallBack"
  />

</template>

<script>
import axios from 'axios'

export default {
  data () {//大转盘样式
    return {
      prizes: [],
      defaultStyle: {
        fontColor: '#ffffff',
        fontSize: '14px'
      },
      blocks: [

          {
            padding: '50px',
            imgs: [{
              src: require('./img/zhuanpan1.png'),
              width: '100%',
              rotate: true
            }]
          }
      ],
      buttons: [

        {
          radius: '35px', background: '#ffdea0',
          imgs: [{ src: require('./img/button2.png'), width: '160%', top: '-185%' }]
        }
      ],
    }
  },
  mounted () {
    {this.getPrizesList()}

  },



  methods: {

    getPrizesList () {
      const prizes = []
      let data = ['1元红包', '100元红包', '0.5元红包', '2元红包', '10元红包', '50元红包']
      data.forEach((item, index) => {
        prizes.push({
          name: item,
          background: index % 2 ? '#9f5ab5' : '#f8e448',
          fonts: [{ text: item, top: '10%' }],
          imgs:[],
        })
      })
      this.prizes = prizes

      {//奖品key，名称，数量的传递
        let prizeList = {
          PrizeID: '',
          PrizeName: '',
          PrizeAmount: ''
        }
        data.forEach((item, index) => {
          prizeList[index] = {
            PrizeId: index,
            PrizeName: item,
            PrizeAmount: 999
          }
          axios.post('http://localhost:8082/api/prizeList', prizeList[index]).then((res) => {
            console.log(res);
          })
        })
        console.log(prizeList)
      }

    },
    startCallBack () {
      function myRand()
      {
        var rand = Math.random()
        if (rand < .3) return 0
        else if (rand < .35) return 1;
        else if (rand < .65) return 2;
        else if (rand < .85) return 3;
        else if (rand < .95) return 4;
        else return 5
      }
      this.$refs.LuckyWheel.play()
      setTimeout(() => {
        this.$refs.LuckyWheel.stop(//设置概率
          myRand()//设置中奖概率
        )
      }, 5000)
    }
    ,

    endCallBack (prize) {
      {//中奖信息传递
        var wheelPrize = {
          prizename: prize.name
        }
        console.log(wheelPrize)
        axios.post('http://localhost:8082/api/wheel', wheelPrize).then((res) => {
          console.log(res);
        });
      }
      alert(`恭喜你获得${prize.name}`)

    },

  }
  }


</script>

