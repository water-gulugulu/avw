<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="区块编号">
          <el-input placeholder="搜索条件" v-model="searchInfo.blockNumber"></el-input>
        </el-form-item>    
        <el-form-item label="合约地址">
          <el-input placeholder="搜索条件" v-model="searchInfo.contract"></el-input>
        </el-form-item>    
        <el-form-item label="发起地址">
          <el-input placeholder="搜索条件" v-model="searchInfo.form"></el-input>
        </el-form-item>    
        <el-form-item label="接收地址">
          <el-input placeholder="搜索条件" v-model="searchInfo.to"></el-input>
        </el-form-item>                
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <!-- <el-form-item>
          <el-button @click="openDialog" type="primary">新增avfEventLog表</el-button>
        </el-form-item> -->
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="区块编号" prop="blockNumber" width="120"></el-table-column> 
    
    <el-table-column label="合约地址" prop="contract" width="120"></el-table-column> 
    
    <el-table-column label="发起地址" prop="form" width="120"></el-table-column> 
    
    <el-table-column label="接收地址" prop="to" width="120"></el-table-column> 
    
    <el-table-column label="日志索引" prop="index" width="120"></el-table-column> 
    
    <el-table-column label="事务索引" prop="txIndex" width="120"></el-table-column> 
    
    <el-table-column label="事件名称" prop="name" width="120"></el-table-column> 
    
    <el-table-column label="操作金额" prop="number" width="120"></el-table-column> 
    
    <el-table-column label="tokens" prop="tokens" width="120"></el-table-column> 
    
    <el-table-column label="时间戳" prop="createTime" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateAvfEventLog(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="区块编号:">
            <el-input v-model="formData.blockNumber" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="合约地址:">
            <el-input v-model="formData.contract" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="发起地址:">
            <el-input v-model="formData.form" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="接收地址:">
            <el-input v-model="formData.to" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="日志索引:">
            <el-input v-model="formData.index" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="事务索引:">
            <el-input v-model="formData.txIndex" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="事件名称:">
            <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="操作金额:"><el-input v-model.number="formData.number" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="tokens:">
            <el-input v-model="formData.tokens" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="时间戳:"><el-input v-model.number="formData.createTime" clearable placeholder="请输入"></el-input>
      </el-form-item>
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createAvfEventLog,
    deleteAvfEventLog,
    deleteAvfEventLogByIds,
    updateAvfEventLog,
    findAvfEventLog,
    getAvfEventLogList
} from "@/api/avfEventLog";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "AvfEventLog",
  mixins: [infoList],
  data() {
    return {
      listApi: getAvfEventLogList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            blockNumber:"",
            contract:"",
            form:"",
            to:"",
            index:"",
            txIndex:"",
            name:"",
            number:0,
            tokens:"",
            createTime:0,
            
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10              
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      deleteRow(row){
        this.$confirm('确定要删除吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
           this.deleteAvfEventLog(row);
        });
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteAvfEventLogByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          if (this.tableData.length == ids.length) {
              this.page--;
          }
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateAvfEventLog(row) {
      const res = await findAvfEventLog({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reavfEventLog;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          blockNumber:"",
          contract:"",
          form:"",
          to:"",
          index:"",
          txIndex:"",
          name:"",
          number:0,
          tokens:"",
          createTime:0,
          
      };
    },
    async deleteAvfEventLog(row) {
      const res = await deleteAvfEventLog({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        if (this.tableData.length == 1) {
            this.page--;
        }
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createAvfEventLog(this.formData);
          break;
        case "update":
          res = await updateAvfEventLog(this.formData);
          break;
        default:
          res = await createAvfEventLog(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>
