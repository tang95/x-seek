import {LoginForm} from '@ant-design/pro-components';
import React from "react";
import styles from './index.less';
import {Helmet, useRequest} from '@umijs/max';
import {Divider, Flex, Spin} from 'antd';
import {AuthService} from "@/services";
import Github from "./component/Github";
import DingTalk from "./component/DingTalk";

const LoginPage: React.FC = () => {
    const {data, loading} = useRequest(AuthService.providers, {
        formatResult: (data) => {
            return data
        }
    });

    const methods = () => {
        if (loading) {
            return (
                <Flex justify={"center"} align={"center"} gap={"middle"}>
                    <Spin/>
                </Flex>
            )
        }
        if (!data) {
            return null
        }
        return (
            <Flex justify={"center"} align={"center"} gap={"middle"}>
                <Github providers={data}/>
                <DingTalk providers={data}/>
            </Flex>
        )
    }
    return (
        <div className={styles.container}>
            <Helmet><title>登录 - X-Seek</title></Helmet>
            <LoginForm
                logo="/favicon.png"
                title="X-Seek"
                subTitle="AI 驱动的事件管理"
                contentStyle={{
                    minWidth: 280,
                    justifyContent: "center",
                    maxWidth: '75vw',
                }}
                containerStyle={{justifyContent: "center"}}
                submitter={false}
                actions={[
                    <div>
                        <Divider plain>
                            登录方式
                        </Divider>
                        {methods()}
                    </div>
                ]}
            >
            </LoginForm>
        </div>
    );
};

export default LoginPage;
