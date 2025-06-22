import { Outlet } from "react-router-dom";
import { ContentContainer, MainContainer } from "./styled";
import { Flex, Typography } from "antd";

export const AuthLayout = () => {
  return (
    <MainContainer>
      {/* <Layout
      // position={"absolute"} left={"60px"} top={"10px"}
      > */}
      <Flex gap="middle" vertical>
        {/* <img width={40} height={40} src={logo} alt="" /> */}
        <Typography>e-Report</Typography>
      </Flex>
      {/* </Layout> */}
      <ContentContainer>
        <Outlet />
      </ContentContainer>
    </MainContainer>
  );
};
