<script setup lang="ts">
import { ref } from 'vue'

const toggleList = (index:number) => {
    lists.value[index].isExpanded = !lists.value[index].isExpanded;
};



const lists = ref([
    {
        title: "杂烩",
        isExpanded: false,
        items: [
            "海绵宝宝",
        ]
    },
    {
        title: "粉丝社区",
        isExpanded: false,
        items: [
            "ikun",
        ]
    },
    {
        title: "搞怪",
        isExpanded: false,
        items: [
            "猥琐萌",
        ]
    }
]);




</script>

<template>
    <div class="container-body">

        <!-- 侧边栏 -->
        <div class="sidebar">
            <div class="components-container">

                <div class="expandable-list" v-for="(list, index) in lists" :key="index">
                    <div class="list-header" @click="()=>toggleList(index)">
                        <div class="list-title">{{ list.title }}</div>
                        <div class="list-icon" :class="{ expanded: list.isExpanded }">▼</div>
                    </div>
                    <div class="list-content" :class="{ expanded: list.isExpanded }">
                        <div class="list-items">
                            <div class="list-item" v-for="(item, index) in list.items" :key="index">
                                <div class="item-icon">●</div>
                                <div class="item-text">{{ item }}</div>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </div>

        <div class="container">
            我是很多的图片
        </div>

    </div>
</template>

<style lang="less" scoped>

.container-body{
    background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            padding: 20px;
}
.container {
    width: 100%;
    max-width: 800px;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 15px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
    overflow: hidden;
    border: 1px solid rgba(255, 255, 255, 0.2);
    padding: 30px;
}


h1 {
    color: white;
    font-size: 2.5rem;
    margin-bottom: 10px;
    text-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
}

.subtitle {
    color: rgba(255, 255, 255, 0.8);
    font-size: 1.2rem;
}

.components-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

/* 可展开列表组件样式 */
.expandable-list {
    background: rgba(255, 255, 255, 0.15);
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
    border: 1px solid rgba(255, 255, 255, 0.2);
}

.expandable-list:hover {
    transform: translateY(-3px);
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.2);
}

.list-header {
    padding: 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    cursor: pointer;
    background: rgba(0, 0, 0, 0.1);
    transition: background 0.3s ease;
}

.list-header:hover {
    background: rgba(0, 0, 0, 0.15);
}

.list-title {
    color: white;
    font-size: 1.4rem;
    font-weight: 600;
}

.list-icon {
    color: white;
    transition: transform 0.3s ease;
    font-size: 1.5rem;
}

.list-icon.expanded {
    transform: rotate(180deg);
}

.list-content {
    overflow: hidden;
    max-height: 0;
    transition: max-height 0.5s ease;
}

.list-content.expanded {
    max-height: 500px;
}

.list-items {
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.list-item {
    padding: 15px;
    background: rgba(255, 255, 255, 0.1);
    border-radius: 8px;
    color: white;
    display: flex;
    align-items: center;
    gap: 12px;
    transition: all 0.3s ease;
    border-left: 4px solid transparent;
}

.list-item:hover {
    background: rgba(255, 255, 255, 0.2);
    transform: translateX(5px);
    border-left-color: #6a11cb;
}

.item-icon {
    font-size: 1.2rem;
    color: #6a11cb;
}

.item-text {
    flex: 1;
}



/* 响应式设计 */
@media (max-width: 768px) {
    .container {
        padding: 20px;
    }

    h1 {
        font-size: 2rem;
    }

    .list-title {
        font-size: 1.2rem;
    }
}
</style>