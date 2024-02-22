<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="类型" prop="inOut">
          <el-input v-model.number="searchInfo.inOut" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="分类" prop="categoryId">
          <el-input
            v-model.number="searchInfo.categoryId"
            placeholder="搜索条件"
          />
        </el-form-item>

        <el-form-item label="发生时间" prop="occurTime">
          <div class="block">
            <span class="demonstration">
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（包含）"
              >
              </el-tooltip>
            </span>
            <el-date-picker
              v-model="value2"
              type="monthrange"
              unlink-panels
              range-separator="To"
              start-placeholder="开始年月"
              end-placeholder="结束年月"
              :shortcuts="shortcuts"
            />
          </div>
          <!-- <el-date-picker v-model="searchInfo.startOccurTime" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endOccurTime ? time.getTime() > searchInfo.endOccurTime.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endOccurTime" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startOccurTime ? time.getTime() < searchInfo.startOccurTime.getTime() : false"></el-date-picker> -->
        </el-form-item>

        <el-form-item label="备注" prop="remark">
          <el-input v-model="searchInfo.remark" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-popover v-model="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link @click="deleteVisible = false"
              >取消</el-button
            >
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
              >删除</el-button
            >
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{
            formatDate(scope.row.createdAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="类型;1：收入（钱进） 2：支出（钱出）"
          prop="inOut"
          width="120"
        />
        <el-table-column
          align="left"
          label="分类;1:存取款"
          prop="categoryId"
          width="120"
        />
        <el-table-column align="left" label="发生时间" width="180">
          <template #default="scope">{{
            formatDate(scope.row.occurTime)
          }}</template>
        </el-table-column>
        <el-table-column align="left" label="金额" prop="money" width="120" />
        <el-table-column align="left" label="备注" prop="remark" width="120" />
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateMoneyAccountRecordFunc(scope.row)"
              >修改</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close
    >
      <el-form
        :model="formData"
        label-position="right"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="类型;1：收入（钱进） 2：支出（钱出）:" prop="type">
          <el-input
            v-model.number="formData.inOut"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="分类;1:存取款:" prop="categoryId">
          <el-input
            v-model.number="formData.categoryId"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="发生时间:" prop="occurTime">
          <el-date-picker
            v-model="formData.occurTime"
            type="date"
            style="width: 100%"
            placeholder="选择日期"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="金额:" prop="money">
          <el-input-number
            v-model="formData.money"
            style="width: 100%"
            :precision="2"
            :clearable="true"
          />
        </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input
            v-model="formData.remark"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "MoneyAccountRecord",
};
</script>

<script setup>
import {
  createMoneyAccountRecord,
  deleteMoneyAccountRecord,
  deleteMoneyAccountRecordByIds,
  updateMoneyAccountRecord,
  findMoneyAccountRecord,
  getMoneyAccountRecordList,
} from "@/api/moneyAccountRecord";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  inOut: 0,
  categoryId: 0,
  occurTime: new Date(),
  money: 0,
  remark: "",
});

// 验证规则
const rule = reactive({
  inOut: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  categoryId: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  occurTime: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  money: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
});

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
  occurTime: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startOccurTime && !searchInfo.value.endOccurTime) {
          callback(new Error("请填写结束日期"));
        } else if (
          !searchInfo.value.startOccurTime &&
          searchInfo.value.endOccurTime
        ) {
          callback(new Error("请填写开始日期"));
        } else if (
          searchInfo.value.startOccurTime &&
          searchInfo.value.endOccurTime &&
          (searchInfo.value.startOccurTime.getTime() ===
            searchInfo.value.endOccurTime.getTime() ||
            searchInfo.value.startOccurTime.getTime() >
              searchInfo.value.endOccurTime.getTime())
        ) {
          callback(new Error("开始日期应当早于结束日期"));
        } else {
          callback();
        }
      },
      trigger: "change",
    },
  ],
});

const elFormRef = ref();
const elSearchFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const searchInfo = ref({});

// 发生时间
const shortcuts = [
  {
    text: "本月",
    value: [new Date(), new Date()],
  },
  {
    text: "近六月",
    value: () => {
      const end = new Date();
      const start = new Date();
      start.setMonth(start.getMonth() - 6);
      return [start, end];
    },
  },
  {
    text: "今年",
    value: () => {
      const end = new Date();
      const start = new Date(new Date().getFullYear(), 0);
      return [start, end];
    },
  },
  {
    text: "近一年",
    value: () => {
      const end = new Date();
      const start = new Date();
      start.setMonth(start.getMonth() - 6);
      return [start, end];
    },
  },
  {
    text: "近三年",
    value: () => {
      const end = new Date();
      const start = new Date();
      start.setMonth(start.getMonth() - 6);
      return [start, end];
    },
  },
];

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    page.value = 1;
    pageSize.value = 10;
    getTableData();
  });
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getMoneyAccountRecordList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteMoneyAccountRecordFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);

// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.id);
    });
  const res = await deleteMoneyAccountRecordByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");

// 更新行
const updateMoneyAccountRecordFunc = async (row) => {
  const res = await findMoneyAccountRecord({ id: row.id });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.remoneyAccountRecord;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteMoneyAccountRecordFunc = async (row) => {
  const res = await deleteMoneyAccountRecord({ id: row.id });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false;
  formData.value = {
    inOut: 0,
    categoryId: 0,
    occurTime: new Date(),
    money: 0,
    remark: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createMoneyAccountRecord(formData.value);
        break;
      case "update":
        res = await updateMoneyAccountRecord(formData.value);
        break;
      default:
        res = await createMoneyAccountRecord(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};
</script>

<style scoped>
.demo-date-picker {
  display: flex;
  width: 100%;
  padding: 0;
  flex-wrap: wrap;
}
.demo-date-picker .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  flex: 1;
}
.demo-date-picker .block:last-child {
  border-right: none;
}
.demo-date-picker .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  margin-bottom: 20px;
}
</style>
