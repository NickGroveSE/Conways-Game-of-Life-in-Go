<script setup>
import {reactive, onMounted} from 'vue'
import {Start, Next, Reset, StoreCell} from '../../wailsjs/go/main/App'
import {EventsOn} from '../../wailsjs/runtime/runtime'

// EventsOn("CountUp", (count) => data.resultText = count)

// function countUp() {
//   CountUp().then(result => {
//     data.resultText = result
//   })
// }

  const data = reactive({
    vueCanvas: null,
    iterationsCount: "Iterations: 0",
  })

  onMounted(() => {
    var c = document.getElementById("grid");
    var ctx = c.getContext("2d");    
    data.vueCanvas = ctx;
    drawBoard()
  })

  EventsOn("Start", (filledCells) => periodicFillCallback(iterations, filledCells))
  EventsOn("Next", (filledCells) => singularFillCallback(filledCells))
  EventsOn("Reset", (filledCells) => singularFillCallback(filledCells))

  function start() {
    Start()
  }

  function next() {
    Next().then(result => {
      data.iterationsCount = result
    })
  }

  function reset() {
    Reset().then(result => {
      data.iterationsCount = result
    })
  }

  function periodicFillCallback(iterations, filledCells) {

  }

  function singularFillCallback(filledCells) {
    data.vueCanvas.clearRect(0, 0, 1000, 800)
    drawBoard()

    for (var i = 0; i < filledCells.length; i++) {
      fillCell("#9C27B0", filledCells[i][1] * 20 + 1, filledCells[i][0] * 20 + 1, 19, 19)
    }
  }

  // function start() {

  // }

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
    
    if(cellColorData.data[0] != 0 && cellColorData.data[0] == 156) {
      // data.vueCanvas.fillStyle = "#FFFFFF"
      // data.vueCanvas.fillRect(cellXOrigin + 1, cellYOrigin + 1, 19, 19)
      fillCell("#FFFFFF", cellXOrigin + 1, cellYOrigin + 1, 19, 19)
      StoreCell(cellXPosition, cellYPosition, false)
    } else {
      // data.vueCanvas.fillStyle = "#9C27B0"
      // data.vueCanvas.fillRect(cellXOrigin + 1, cellYOrigin + 1, 19, 19)
      fillCell("#9C27B0", cellXOrigin + 1, cellYOrigin + 1, 19, 19)
      StoreCell(cellXPosition, cellYPosition, true)
    }

  }

  function fillCell(color, x, y, width, height) {
    data.vueCanvas.fillStyle = color
    data.vueCanvas.fillRect(x, y, width, height)
  }

</script>

<template>
  <main>
    <canvas id="grid" width="1000px" height="800px" @click="cellClicked"></canvas>
    <div id="iterations-count">{{ data.iterationsCount }}</div>
    <button class="control-button" id="next-button" @click="next">Next</button>
    <button class="control-button" id="clear-button" @click="reset">Reset Board</button>
    <!-- <button class="control-button" id="start-button" @click="start">Start</button> -->
    
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

canvas {
  background-color: #FFFFFF;
}

#controls {
  margin-top: -6px;
  padding: 20px;
  border-top: 1px solid #a3a2a2;
}

#iterations-count {
  position: absolute;
  top: 15px;
  left: 15px;
  padding: 5px 15px;
  background-color: white;
  border: 1px solid #a3a2a2;
  border-radius: 15px;
}

.control-button {
  position: absolute;
  bottom: 15px;
  padding: 5px 15px;
  background-color: white;
  border: 1px solid #a3a2a2;
  border-radius: 15px;
  font-size: 16px;
  box-shadow: 2px 2px #C187D1;
  transition: 0.5s;
}

#next-button {
  left: 15px;
}

#clear-button {
  left: 90px;
}


#next-button:hover {
  box-shadow: 5px 5px #C187D1;
  left: 12px;
  bottom: 18px;
}

#clear-button:hover {
  box-shadow: 5px 5px #C187D1;
  left: 87px;
  bottom: 18px;
}

</style>
