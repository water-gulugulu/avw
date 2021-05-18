<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增用户账单</el-button>
        </el-form-item>
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
    
    <el-table-column label="用户ID" prop="uid" width="120"></el-table-column> 
    
    <el-table-column label="卡牌ID" prop="cardId" width="120"></el-table-column> 
    
    <el-table-column label="钱包地址" prop="address" width="120"></el-table-column> 
    
    <el-table-column label="账单类型" prop="type" width="120">
         <template slot-scope="scope">{{scope.row.type|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="金额" prop="money" width="120"></el-table-column> 
    
    <el-table-column label="手续费" prop="fees" width="120"></el-table-column> 
    
    <el-table-column label="余额" prop="balance" width="120"></el-table-column> 
    
    <el-table-column label="收入支出" prop="payment" width="120">
         <template slot-scope="scope">{{scope.row.payment|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="支付方式" prop="payType" width="120">
         <template slot-scope="scope">{{scope.row.payType|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="描述" prop="detail" width="120"></el-table-column> 
    
    <el-table-column label="创建时间" prop="createTime" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateAvfUserBill(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
         <el-form-item label="用户ID:"><el-input v-model.number="formData.uid" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="卡牌ID:"><el-input v-model.number="formData.cardId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="钱包地址:">
            <el-input v-model="formData.address" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="账单类型:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.type" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="金额:"><el-input v-model.number="formData.money" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="手续费:"><el-input v-model.number="formData.fees" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="余额:"><el-input v-model.number="formData.balance" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="收入支出:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.payment" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="支付方式:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.payType" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="描述:">
            <el-input v-model="formData.detail" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="创建时间:"><el-input v-model.number="formData.createTime" clearable placeholder="请输入"></el-input>
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
    createAvfUserBill,
    deleteAvfUserBill,
    deleteAvfUserBillByIds,
    updateAvfUserBill,
    findAvfUserBill,
    getAvfUserBillList
} from "@/api/avf_user_bill";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "AvfUserBill",
  mixins: [infoList],
  data() {
    return {
      listApi: getAvfUserBillList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            uid:0,
            cardId:0,
            address:"",
            type:false,
            money:0,
            fees:0,
            balance:0,
            payment:false,
            payType:false,
            detail:"",
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
        if (this.searchInfo.type==""){
          this.searchInfo.type=null
        }           
        if (this.searchInfo.payment==""){
          this.searchInfo.payment=null
        }        
        if (this.searchInfo.payType==""){
          this.searchInfo.payType=null
        }        
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
           this.deleteAvfUserBill(row);
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
        const res = await deleteAvfUserBillByIds({ ids })
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
    async updateAvfUserBill(row) {
      const res = await findAvfUserBill({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reavfUserBill;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          uid:0,
          cardId:0,
          address:"",
          type:false,
          money:0,
          fees:0,
          balance:0,
          payment:false,
          payType:false,
          detail:"",
          createTime:0,
          
      };
    },
    async deleteAvfUserBill(row) {
      const res = await deleteAvfUserBill({ ID: row.ID });
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
          res = await createAvfUserBill(this.formData);
          break;
        case "update":
          res = await updateAvfUserBill(this.formData);
          break;
        default:
          res = await createAvfUserBill(this.formData);
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
