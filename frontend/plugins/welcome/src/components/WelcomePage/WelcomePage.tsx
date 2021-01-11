import React, { useEffect, FC } from 'react';
import { Button,Typography, Grid, Link } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles,Theme ,createStyles} from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Avatar from '@material-ui/core/Avatar';
import {  Cookies  } from 'react-cookie/cjs';//cookie
import { useCookies } from 'react-cookie/cjs';//cookie


import { Image1Base64Function } from '../../image/Image1';
import { Image2Base64Function } from '../../image/Image2';
import { Image3Base64Function } from '../../image/Image3';
import { Image4Base64Function } from '../../image/Image4';
import { Image5Base64Function } from '../../image/image5_home';
import { Image6Base64Function } from '../../image/image6_home';

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles((theme: Theme) =>
createStyles({
  root: {
    maxWidth: 325,
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  
}),
);

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
  imgsut: string;
  linkto: string;
};


export function CardTeam({ name, id, system ,imgsut,linkto}: ProfileProps) {
  
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
        <Link
        href = {linkto}
        >
          <CardMedia
            component="img"
            height="140"
            image={imgsut}  
          />
           </Link>
          <CardContent>
            <Typography gutterBottom variant="subtitle1" component="h2">
              <br/>{system}
              <br/>{id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const WelcomePage: FC<{}> = () => {
  const [cookies, setCookie, removeCookie] = useCookies(['cookiename']);

  const Logout = async () => {

    removeCookie('Name', { path: '/' })
    removeCookie('Password', { path: '/' })
    removeCookie('Log', { path: '/' })
    removeCookie('Status', { path: '/' })

  }
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบผู้ป่อยนอก`}>
      {/*
      <Link href = "/">
      <Button
                onClick={() => {
                  Logout();
                }}
                variant="contained"
              >
                 Logout
             </Button>
             </Link>
             */}
      </Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Grid container>
        
          <CardTeam name={"นางสาวณัชพร กลิ่นรอด"} id={"B6102845"} system={"ระบบการเงิน"} imgsut = {Image1Base64Function} linkto = "/createBill"></CardTeam>
          <CardTeam name={"นางสาวพรรวินท์ กุดแถลง"} id={"B6103217"} system={"ระบบลงทะเบียนผู้ป่วยนอก"} imgsut = {Image2Base64Function}linkto = "/createPatientrecord"></CardTeam>
          <CardTeam name={"นายคฑาเดช เขียนชัยนาจ"} id={"B6103866"} system={"ระบบสิทธิ์ผู้ป่วย"} imgsut = {Image3Base64Function}linkto = "/create_Patientrights"></CardTeam>
          <CardTeam name={"นายวัชรพงษ์ ทาระมล"} id={"B6107505"} system={"ระบบการรักษาผู้ป่วย"} imgsut = {Image4Base64Function}linkto = "/createTreatment"></CardTeam>
          <CardTeam name={"นางสาวปอรรัชม์ ปานใจนาม"} id={"B6109868"} system={"ระบบซักประวัติ"} imgsut = {Image5Base64Function}linkto = "/Historytaking"></CardTeam>
          <CardTeam name={"นายธนวรรต สีแก้วสิ่ว"} id={"B6115586"} system={"ระบบข้อมูลแพทย์"} imgsut = {Image6Base64Function}linkto = "/Doctorinfo"></CardTeam>
         
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
