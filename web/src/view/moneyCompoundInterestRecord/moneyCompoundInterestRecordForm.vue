<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="存款时间:" prop="depositTime">
          <el-date-picker v-model="formData.depositTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="到期时间:" prop="expirationTime">
          <el-date-picker v-model="formData.expirationTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="本金:" prop="principal">
          <el-input-number v-model="formData.principal" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="利息:" prop="interest">
          <el-input-number v-model="formData.interest" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="本息和:" prop="principalInterest">
          <el-input-number v-model="formData.principalInterest" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MoneyCompoundInterestRecord'
}
</script>

<script setup>
import {
  createMoneyCompoundInterestRecord,
  updateMoneyCompoundInterestRecord,
  findMoneyCompoundInterestRecord
} from '@/api/moneyCompoundInterestRecord'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            depositTime: new Date(),
            expirationTime: null,
            principal: null,
            interest: null,
            principalInterest: null,
        })
// 验证规则
const rule = reactive({
               principal : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据id 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findMoneyCompoundInterestRecord({ id: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.remoneyCompoundInterestRecord
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createMoneyCompoundInterestRecord(formData.value)
               break
             case 'update':
               res = await updateMoneyCompoundInterestRecord(formData.value)
               break
             default:
               res = await createMoneyCompoundInterestRecord(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
