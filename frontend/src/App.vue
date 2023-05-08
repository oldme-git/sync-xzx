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
                <button class="input-btn" @click="xzxConn">连接测试</button>
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
                <button class="input-btn" @click="dbConn">连接测试</button>
                <button class="input-btn" @click="saveConf">保存配置</button>
            </div>
        </div>
        <div class="card">
            <h1>其他信息</h1>
            <p>上次同步时间：</p>
            <p>{{ syncTime.prev }}</p>
            <p>下次同步时间：</p>
            <p>{{ syncTime.next }}</p>
            <p>同步状态：</p>
            <p>已经同步 {{ progress.num }} 条数据</p>
            <div class="input-group">
                <input type="text" class="input-field" v-model="conf.ext.syncTime" :readonly="disabledAutoSyncBtn">
                <label class="input-label">自动同步时间间隔（分）</label>
            </div>
            <div class="input-group">
                <button class="input-btn" @click="saveConf">保存配置</button>
                <button class="input-btn" @click="sync" :disabled="disabledSyncBtn">立即同步</button>
                <button class="input-btn" @click="cleanLog">清空日志信息</button>
                <br>
                <button class="input-btn" @click="startAutoSync" :disabled="disabledAutoSyncBtn">开启自动同步</button>
                <button class="input-btn" @click="closeAutoSync" :disabled="!disabledAutoSyncBtn">关闭自动同步</button>
            </div>
        </div>
        <div class="card log">
            <h1>日志信息</h1>
            <p v-for="item in logList">{{ item }}</p>
        </div>
    </div>
</template>

<script setup>
import {
    SaveConf as _sc,
    ReadConf as _rf,
    XzxConnTest as _xct,
    DbConnTest as _dct,
    Sync as _sync,
    StartAutoSync as _sas,
    CloseAutoSync as _cas
} from "../wailsjs/go/main/App.js"
import {watch, onMounted, ref} from "vue"
import {EventsOn} from "../wailsjs/runtime/runtime.js"

const conf = ref({
    xzx: {
        ip: "172.20.2.220",
        port: "8500",
        terminal: "10",
        sysCode: "21",
        timeout: "20",
    },
    db: {
        ip: "127.0.0.1",
        port: "3306",
        dbName: "traceint",
        user: "root",
        pass: "123456",
    },
    ext: {
        syncTime: "180",
        isAutoSync: "0"
    },
})
const logList = ref([])
const progress = ref({
    num: 0,
    show: false
})
const disabledSyncBtn = ref(false)
const disabledAutoSyncBtn = ref(false)
const syncTime = ref({
    prev: "",
    next: "",
})

// 定时清除日志
setInterval(() => {
    logList.value = []
}, 3600 * 1000)

// 同步进度
EventsOn("progress", (res) => {
    // 开启进度显示
    progress.value.show = true
    progress.value.num = 0
    progress.value.num = res
})

// 自动同步时日志
EventsOn("autoSyncStatus", (res) => {
    pushLog(res)
})

// 记录同步时间
EventsOn("autoTime", (res) => {
    try {
        syncTime.value = JSON.parse(res)
    } catch (e) {}
})

onMounted(() => {
    readConf()
})

// 读取配置
function readConf() {
    _rf().then((res) => {
        try {
            conf.value = JSON.parse(res)
            // 初始化自动同步
            if (conf.value.ext.isAutoSync === "1") {
                startAutoSync()
                changeAutoSyncBtn()
            }
        } catch (e) {}
    })
}

// 保存配置
function saveConf() {
    const jsonConf = JSON.stringify(conf.value)
    _sc(jsonConf).then((res) => {
        pushLog(res)
    })
}

// 新中新连接测试
function xzxConn() {
    _xct().then((res) => {
        pushLog(res)
    })
}

// 数据库连接测试
function dbConn() {
    _dct().then((res) => {
        pushLog(res)
    })
}

// 立即同步
function sync() {
    // 禁止重复点击
    disabledSyncBtn.value = true

    _sync().then((res) => {
        pushLog(res)
        // 解除限制
        disabledSyncBtn.value = false
    })
}

// 开启自动同步
function startAutoSync() {
    conf.value.ext.isAutoSync = "1"
    _sas().then((res) => {
        pushLog(res)
    })
}

// 关闭自动同步
function closeAutoSync() {
    conf.value.ext.isAutoSync = "0"
    _cas().then((res) => {
        pushLog(res)
    })
}

// 清空日志
function cleanLog() {
    logList.value = []
}

// 增加一个日志
function pushLog(text) {
    let myDate = new Date()
    let hours =  String(myDate.getHours()).padStart(2, "0");
    let minutes = String(myDate.getMinutes()).padStart(2, "0");
    let seconds = String(myDate.getSeconds()).padStart(2, "0");
    let now = hours + ":" + minutes + ":" + seconds
    logList.value.push(now + " " + text)
}

// 改变自动同步按钮状态
function changeAutoSyncBtn() {
    disabledAutoSyncBtn.value = conf.value.ext.isAutoSync === "1";
}

// 监听自动同步状态
watch(() => conf.value.ext.isAutoSync, () => {
    changeAutoSyncBtn()
}, { deep: true })

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
    width: 50%;
    height: 320px;
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
    width: 70%;
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

.input-field[readonly] {
    background-color: #eeeeee;
}

.input-label {
    position: absolute;
    top: 0;
    left: 0;
    font-size: 14px;
    pointer-events: none;
    transition: all 0.3s ease-out;
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

.input-btn[disabled] {
    background-color: #2972a2;
    box-shadow: none;
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
