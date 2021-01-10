import React, { useEffect, FC } from 'react';
import { Typography, Grid } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Button,
} from '@backstage/core';
import { makeStyles,Theme ,createStyles} from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Avatar from '@material-ui/core/Avatar';

import {  Cookies  } from 'react-cookie/cjs';//cookie

import { Image1Base64Function } from '../../image/Image1';
import { Image2Base64Function } from '../../image/Image2';
import { Image3Base64Function } from '../../image/Image3';
import { Image4Base64Function } from '../../image/Image4';
import { Image5Base64Function } from '../../image/Image5';
import { Image6Base64Function } from '../../image/Image6';

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles((theme: Theme) =>
createStyles({
  root: {
    maxWidth: 325,
  },
  button: {
    display: 'block',
    marginTop: theme.spacing(2),
  },
  
}),
);
const styles = 
{
media: {
  height: 0,
  paddingTop: '20%', // 16:9,
  marginTop:'30'
}
  };
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
          <CardMedia
            component="img"
            height="140"
            image={imgsut}
            
          />
          <CardContent>
            <Typography gutterBottom variant="subtitle1" component="h2">
            <Button
            to = {linkto}
            variant = "outlined"
            >
              {system}
            </Button>
              <br/>{id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const WelcomePage: FC<{}> = () => {
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบผู้ป่อยนอก`}></Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Grid container>
        
          <CardTeam name={"นางสาวณัชพร กลิ่นรอด"} id={"B6102845"} system={"ระบบการเงิน"} imgsut = {Image1Base64Function} linkto = "/createbill"></CardTeam>

          <CardTeam name={"นางสาวพรรวินท์ กุดแถลง"} id={"B6103217"} system={"ระบบลงทะเบียนผู้ป่วยนอก"} imgsut = {Image2Base64Function}linkto = "/createPatientrecord"></CardTeam>
          <CardTeam name={"นายคฑาเดช เขียนชัยนาจ"} id={"B6103866"} system={"ระบบสิทธิ์ผู้ป่วย"} imgsut = {Image3Base64Function}linkto = "/create_Patientrights"></CardTeam>
          <CardTeam name={"นายวัชรพงษ์ ทาระมล"} id={"B6107505"} system={"ระบบการรักษาผู้ป่วย"} imgsut = {Image4Base64Function}linkto = "/createTreatment"></CardTeam>
          <CardTeam name={"นางสาวปอรรัชม์ ปานใจนาม"} id={"B6109868"} system={"ระบบซักประวัติ"} imgsut = {Image5Base64Function}linkto = "/createHistorytaking"></CardTeam>
          <CardTeam name={"นายธนวรรต สีแก้วสิ่ว"} id={"B6115586"} system={"ระบบข้อมูลแพทย์"} imgsut = {Image6Base64Function}linkto = "/Doctorinfo"></CardTeam>
         
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
