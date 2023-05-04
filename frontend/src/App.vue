<template>
    <div class="warp">
        <div class="card">
            <h1>第三方服务器设置</h1>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.xzx.ip">
                <label class="input-label">IP</label>
            </div>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.xzx.port">
                <label class="input-label">端口</label>
            </div>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.xzx.terminal">
                <label class="input-label">站点</label>
            </div>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.xzx.sysCode">
                <label class="input-label">子系统代码</label>
            </div>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.xzx.timeout">
                <label class="input-label">超时时长</label>
            </div>
            <div class="input-group">
                <button class="input-btn">连接测试</button>
                <button class="input-btn" @click="saveConf">保存配置</button>
            </div>
        </div>
        <div class="card">
            <h1>接收数据库配置</h1>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.db.ip">
                <label class="input-label">IP</label>
            </div>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.db.port">
                <label class="input-label">端口</label>
            </div>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.db.dbName">
                <label class="input-label">数据库名</label>
            </div>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.db.user">
                <label class="input-label">账号</label>
            </div>
            <div class="input-group">
                <input type="password" class="input-field" v-model="conf.db.pass">
                <label class="input-label">密码</label>
            </div>
            <div class="input-group">
                <button class="input-btn">连接测试</button>
                <button class="input-btn" @click="saveConf">保存配置</button>
            </div>
        </div>
        <div class="card">
            <h1>其他信息</h1>
            <p>上次同步时间：</p>
            <p>2022-02-03 12:00:00</p>
            <p>下次同步时间：</p>
            <p>2022-02-03 12:00:00</p>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.ext.syncTime">
                <label class="input-label">自动同步时间间隔（分）</label>
            </div>
            <div class="input-group">
                <button class="input-btn" @click="saveConf">保存配置</button>
                <button class="input-btn">立即同步</button>
                <button class="input-btn">开启自动同步</button>
                <button class="input-btn">关闭自动同步</button>
                <button class="input-btn" @click="cleanLog">清空日志信息</button>
            </div>
        </div>
        <div class="card log">
            <h1>日志信息</h1>
            <p v-for="item in logList">{{ item }}</p>
        </div>
    </div>
</template>

<script setup>
import {SaveConf} from "../wailsjs/go/main/App.js"
import {ref} from "vue"

const conf = ref({
    xzx: {
        ip: "172.20.2.220",
        port: 8500,
        terminal: 21,
        sysCode: 10,
        timeout: 20,
    },
    db: {
        ip: "127.0.0.1",
        port: 3600,
        dbName: "traceint",
        user: "root",
        pass: "123456",
    },
    ext: {
        syncTime: 1800,
    },
})
const logList = ref([])

function saveConf() {
    const jsonConf = JSON.stringify(conf.value)
    console.log(jsonConf)
    SaveConf(jsonConf).then((res) => {
        logList.value.push(res)
    })
}

// 清空日志
function cleanLog() {
    logList.value = []
}

</script>

<style scoped>
.warp {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}

.card {
    border: 1px solid #ccc;
    padding: 10px;
    width: 240px;
    height: 330px;
    box-sizing: border-box;
}

.card:nth-child(1) {
    border-right: none;
    border-bottom: none;
}

.card:nth-child(4) {
    border-left: none;
    border-top: none;
}

.card h1 {
    font-size: 16px;
    color: #666;
    margin-top: 0;
    margin-bottom: 6px;
}

.input-group {
    padding-top: 18px;
    width: 200px;
    position: relative;
    margin-bottom: 4px;
}

.input-field {
    font-size: 14px;
    padding: 2px 5px;
    border: 1px solid #ccc;
    width: 100%;
    background: transparent;
}

.input-field:focus {
    outline: none;
    border: 1px solid #3498db;
}

.input-label {
    position: absolute;
    top: 5px;
    font-size: 14px;
    pointer-events: none;
    transition: all 0.3s ease-out;
}

.input-field:focus ~ .input-label,
.input-field:not(:placeholder-shown) ~ .input-label {
    top: 0;
    font-size: 14px;
    color: #3498db;
}

.input-btn {
    display: inline-block;
    border-radius: 2px;
    background-color: #3498db;
    color: #fff;
    font-size: 14px;
    text-align: center;
    text-decoration: none;
    text-transform: uppercase;
    transition: all 0.3s ease;
    border: 1px solid #287bb3;
    margin: 0 10px 10px 0;
}

.input-btn:hover {
    background-color: #3498db;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
}

.card p {
    margin-top: 4px;
    margin-bottom: 4px;
    color: #0f6ba8;
    font-size: 14px;
}

.log {
    overflow: auto;
}
</style>
