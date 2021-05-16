<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">              
        <el-form-item label="事务哈希">
          <el-input placeholder="搜索条件" v-model="searchInfo.txHash"></el-input>
        </el-form-item>    
        <el-form-item label="区块编号">
          <el-input placeholder="搜索条件" v-model="searchInfo.block"></el-input>
        </el-form-item>        
        <el-form-item label="支付地址">
          <el-input placeholder="搜索条件" v-model="searchInfo.from"></el-input>
        </el-form-item>    
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <!--<el-form-item>
          <el-button @click="openDialog" type="primary">新增订单</el-button>
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
    
    <el-table-column label="用户ID" prop="uid" width="120"></el-table-column> 
    
    <el-table-column label="订单编号" prop="orderSn" width="120"></el-table-column> 
    
    <el-table-column label="支付价格" prop="price" width="120"></el-table-column> 
    
    <el-table-column label="购买数量" prop="num" width="120"></el-table-column> 
    
    <el-table-column label="剩余数量" prop="number" width="120"></el-table-column> 
    
    <el-table-column label="订单状态" prop="status" width="120">
         <template slot-scope="scope">{{scope.row.status|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="支付时间" prop="payTime" width="120"></el-table-column> 
    
    <el-table-column label="事务哈希" prop="txHash" width="120"></el-table-column> 
    
    <el-table-column label="区块编号" prop="block" width="120"></el-table-column> 
    
    <el-table-column label="手续费" prop="gas" width="120"></el-table-column> 
    
    <el-table-column label="手续费价格" prop="gasPrice" width="120"></el-table-column> 
    
    <el-table-column label="支付地址" prop="from" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateAvfOrder(scope.row)" size="small" type="primary" icon="el-icon-edit">变更</el-button>
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
       
         <el-form-item label="订单编号:">
            <el-input v-model="formData.orderSn" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="支付价格:">
              <el-input-number v-model="formData.price" :precision="2" clearable></el-input-number>
       </el-form-item>
       
         <el-form-item label="购买数量:"><el-input v-model.number="formData.num" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="剩余数量:"><el-input v-model.number="formData.number" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="订单状态:">
            <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
      </el-form-item>
       
         <el-form-item label="支付时间:"><el-input v-model.number="formData.payTime" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="事务哈希:">
            <el-input v-model="formData.txHash" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="区块编号:">
            <el-input v-model="formData.block" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="手续费:">
            <el-input v-model="formData.gas" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="手续费价格:">
            <el-input v-model="formData.gasPrice" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="支付地址:">
            <el-input v-model="formData.from" clearable placeholder="请输入" ></el-input>
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
    createAvfOrder,
    deleteAvfOrder,
    deleteAvfOrderByIds,
    updateAvfOrder,
    findAvfOrder,
    getAvfOrderList
} from "@/api/avf_order";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "AvfOrder",
  mixins: [infoList],
  data() {
    return {
      listApi: getAvfOrderList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            uid:0,
            orderSn:"",
            price:0,
            num:0,
            number:0,
            status:false,
            payTime:0,
            txHash:"",
            block:"",
            gas:"",
            gasPrice:"",
            from:"",
            
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
        if (this.searchInfo.status==""){
          this.searchInfo.status=null
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
           this.deleteAvfOrder(row);
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
        const res = await deleteAvfOrderByIds({ ids })
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
    async updateAvfOrder(row) {
      const res = await findAvfOrder({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reavfOrder;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          uid:0,
          orderSn:"",
          price:0,
          num:0,
          number:0,
          status:false,
          payTime:0,
          txHash:"",
          block:"",
          gas:"",
          gasPrice:"",
          from:"",
          
      };
    },
    async deleteAvfOrder(row) {
      const res = await deleteAvfOrder({ ID: row.ID });
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
          res = await createAvfOrder(this.formData);
          break;
        case "update":
          res = await updateAvfOrder(this.formData);
          break;
        default:
          res = await createAvfOrder(this.formData);
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
