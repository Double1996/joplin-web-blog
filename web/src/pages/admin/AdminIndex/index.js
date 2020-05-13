import React, {useState} from 'react';
import {Layout, Menu, Breadcrumb, Icon, message} from 'antd';
import {Route} from "react-router-dom";

function AdminIndex(props) {
    const [collapsed, setCollapsed] = useState(false)

    const handleClickArticle = e=>{
        console.log(e.item.props)
        if(e.key=='addArticle'){
            props.history.push('/index/add')
        }else{
            props.history.push('/index/list')
        }
    }
}

return (
    <Layout style={{minHeight: '100vh'}}>
        <Sider collapsible collapsed={collapsed} onCollapse={onCollapse}>
            <div className="logo"> JSPang System</div>
            >
            <Menu theme="dark" defaultSelectedKeys={['1']} mode="inline">
                <Menu.Item key="1">
                    <Icon type="pie-chart"/>
                    <span>工作台</span>
                </Menu.Item>

                <SubMenu
                    key="sub1"
                    onClick={handleClickArticle}
                    title={
                        <span>
                  <Icon type="desktop"/>
                  <span>文章管理</span>
                </span>
                    }
                >
                    <Menu.Item key="addArticle">添加文章</Menu.Item>
                    <Menu.Item key="articleList">文章列表</Menu.Item>
                </SubMenu>

                <Menu.Item key="9">
                    <Icon type="file"/>
                    <span>留言管理</span>
                </Menu.Item>

                <Menu.Item key="10" onClick={handleExit}>
                    <Icon type="file"/>
                    <span>退出登录</span>
                </Menu.Item>
            </Menu>
        </Sider>
        <Layout>

            <Content style={{margin: '0 16px'}}>
                <Breadcrumb style={{margin: '16px 0'}}>
                    <Breadcrumb.Item>后台管理</Breadcrumb.Item>
                    <Breadcrumb.Item>工作台</Breadcrumb.Item>
                </Breadcrumb>
                <div style={{padding: 24, background: '#fff', minHeight: 360}}>
                    <div>
                        <Route path="/index/" exact component={AddArticle}/>
                        <Route path="/index/add/" exact component={AddArticle}/>
                        <Route path="/index/add/:id" exact component={AddArticle}/>
                        <Route path="/index/list/" component={ArticleList}/>
                        <Route path="/index/bbd/" component={BBDList}/>
                    </div>
                </div>
            </Content>
            <Footer style={{textAlign: 'center'}}>double1996.com</Footer>
        </Layout>
    </Layout>
)

export default  AdminIndex