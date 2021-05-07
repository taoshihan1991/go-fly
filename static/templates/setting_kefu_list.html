{{template "header" }}
<div id="app" style="width:100%">
    <template>
        <el-container v-loading.fullscreen.lock="fullscreenLoading">

            <el-main class="mainMain">
                <el-button style="margin-bottom: 10px;" @click="addKefu" type="primary" size="small">添加客服</el-button>
                <el-table
                        :data="kefuList"
                        border
                        style="width: 100%">
                    <el-table-column
                            prop="img"
                            label="客服头像">
                        <template slot-scope="scope">
                            <el-avatar :size="50"><img :src="scope.row.avator"/></el-avatar>
                        </template>
                    </el-table-column>
                    <el-table-column
                            prop="name"
                            label="客服账号">
                    </el-table-column>
                    <el-table-column
                            prop="nickname"
                            label="客服昵称">
                    </el-table-column>
                    <el-table-column
                            prop="role_name"
                            label="角色">
                    </el-table-column>
                    <el-table-column
                            prop="created_at"
                            label="添加时间">
                    </el-table-column>
                    <el-table-column
                            prop="id"
                            label="操作">
                        <template slot-scope="scope">
                            <el-button @click="getKefu(scope.row.id)" type="primary" size="small" plain>编辑</el-button>
                            <el-button @click="deleteKefu(scope.row.id)" type="danger" size="small" plain>删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-main>

        </el-container>
        <el-dialog
                title="客服"
                :visible.sync="kefuDialog"
                width="30%"
                top="0"
                >
            <el-form ref="kefuForm" :model="kefuForm" :rules="rules" label-width="70px">
                <el-form-item label="用户名"  prop="name">
                    <el-input v-model="kefuForm.name"></el-input>
                </el-form-item>
                <el-form-item label="密码"  prop="password">
                    <el-input v-model="kefuForm.password"></el-input>
                </el-form-item>
                <el-form-item label="昵称"  prop="nickname">
                    <el-input v-model="kefuForm.nickname"></el-input>
                </el-form-item>
                <el-form-item label="头像"  prop="avator">
                    <el-select v-model="kefuForm.avator" placeholder="请选择头像">
                        <el-option :label="'头像'+item" :value="'/static/images/'+item+'.jpg'" v-for="item in [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14]" v-bind:key="item">
                            <el-avatar :size="30" :src="'/static/images/'+item+'.jpg'"></el-avatar>
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="角色"  prop="role_id">
                    <el-select v-model="kefuForm.role_id" placeholder="请选择角色">
                        <el-option :label="item.role_name" :value="item.role_id" v-for="item in roleList" v-bind:key="item.role_id">
                        </el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="kefuDialog = false">取 消</el-button>
                <el-button type="primary" @click="submitKefuForm('kefuForm')">确 定</el-button>
              </span>
        </el-dialog>
    </template>

</div>
</body>
{{template "setting_bottom" .}}
