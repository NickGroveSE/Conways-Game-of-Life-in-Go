<script setup>
import {reactive, onMounted} from 'vue'
import {CountUp, StoreCell} from '../../wailsjs/go/main/App'
import {EventsOn} from '../../wailsjs/runtime/runtime'

// EventsOn("CountUp", (count) => data.resultText = count)

// function countUp() {
//   CountUp().then(result => {
//     data.resultText = result
//   })
// }

  const data = reactive({
    vueCanvas: null
  })

  onMounted(() => {
    var c = document.getElementById("grid");
    var ctx = c.getContext("2d");    
    data.vueCanvas = ctx;
    drawBoard()
  })

  // grid width and height
  var bw = 1000;
  var bh = 800;

  function drawBoard(){
    for (var x = 0; x <= bw; x += 20) {
        data.vueCanvas.moveTo(0.5 + x, 0);
      data.vueCanvas.lineTo(0.5 + x, bh);
    }

    for (var x = 0; x <= bh; x += 20) {
        data.vueCanvas.moveTo(0, 0.5 + x);
        data.vueCanvas.lineTo(bw, 0.5 + x);
    }

    data.vueCanvas.strokeStyle = "#a3a2a2";
    data.vueCanvas.stroke();
  }

  function cellClicked(event) {
    var cellXOrigin = event.clientX - (event.clientX % 20) 
    var cellYOrigin = event.clientY - (event.clientY % 20)

    var cellColorData= data.vueCanvas.getImageData(cellXOrigin + 1, cellYOrigin + 1, 19, 19)

    var cellXPosition = cellXOrigin / 20
    var cellYPosition = cellYOrigin / 20
    
    if(cellColorData.data[0] > 0) {
      data.vueCanvas.fillStyle = "#FFFFFF"
      data.vueCanvas.fillRect(cellXOrigin + 1, cellYOrigin + 1, 19, 19)
      StoreCell(cellXPosition, cellYPosition, false).then(result => {
        alert(result)
      })
    } else {
      data.vueCanvas.fillStyle = "#C187D1"
      data.vueCanvas.fillRect(cellXOrigin + 1, cellYOrigin + 1, 19, 19)
      StoreCell(cellXPosition, cellYPosition, true).then(result => {
        alert(result)
      })
    }

  }

</script>

<template>
  <main>
    <canvas id="grid" width="1000px" height="800px" @click="cellClicked"></canvas>
    <!-- <div id="controls">
      <button type="button" id="zoom-in">+</button>
      <button type="button" id="zoom-out">-</button>
      <button type="button" id="move-left"><-</button>
      <button type="button" id="move-right">-></button>
      <button type="button" id="move-up">^</button>
      <button type="button" id="move-down">v</button>
    </div> -->
  </main>
</template>

<style scoped>

main {
  width: 100vw;
  height: 100vh;
}

#controls {
  margin-top: -6px;
  padding: 20px;
  border-top: 1px solid #a3a2a2;
}

</style>
