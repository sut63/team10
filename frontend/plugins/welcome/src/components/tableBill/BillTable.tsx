import React, { useEffect, FC } from 'react';

import { Content, ContentHeader, Header, Page, pageTheme, } from '@backstage/core';

import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Button,Typography } from '@material-ui/core';
import Paper from '@material-ui/core/Paper';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { Avatar } from '@material-ui/core';


import { DefaultApi } from '../../api/apis';
import { EntBill } from '../../api/models/EntBill';
import { EntTreatment } from '../../api';

import { Cookies } from 'react-cookie/cjs';//cookie

  const cookies = new Cookies(); 
  const Name = cookies.get('Name');

  
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    button: {
      display: 'block',
      marginTop: theme.spacing(2),
    },
    formControl: {
      width: 250,
    },
    selectEmpty: {
      marginTop: theme.spacing(1),
    },
    paper: {
      minWidth: 400,
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    table: {
      minWidth: 750,
    },

  }),
);

import { Image1Base64Function } from '../../image/Image1';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

const BillTable: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [loading, setLoading] = React.useState(true);


  const [bills, setbills] = React.useState<EntBill[]>([]);
  const [treatments, setTreatment] = React.useState<EntTreatment[]>([]);

  useEffect(() => {
    const getBill = async () => {
      const res = await http.listBill({offset:0});
      setLoading(false);
      setbills(res);
    };
    const getTreatment = async () => {
      const res = await http.listTreatment({ offset: 0 });
      setLoading(false);
      setTreatment(res);
    };
    getBill();
    getTreatment();
  }, [loading]);
  return (
    
      <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`Financial System`}>
        <Avatar alt="Remy Sharp" src={Image1Base64Function} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
        <Content>
            <ContentHeader title = {`ใบเสร็จรับเงิน (ผู้ป่วยนอก)`}></ContentHeader>
          <Typography align = 'center'>
                  <TableContainer component={Paper} >
                    <Table className={classes.table} aria-label="simple table">
                      <TableHead>
                        <TableRow>
                          <TableCell align="center" >เลขที่ใบเสร็จ</TableCell>
                          <TableCell align="center">เลขที่การรักษา</TableCell>
                          <TableCell align="center">ประเภทการรักษา</TableCell>
                          <TableCell align="center">ผู้ป่วย</TableCell>
                          <TableCell align="center">ค่ารักษา</TableCell>
                          <TableCell align="center">วันที่ชำระ</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {bills.map(item => (treatments.filter(t => t.id == item.edges?.treatment?.id).map(item2 =>(
                          <TableRow key={item.id}>
                            <TableCell align="center">{item.id}</TableCell>
                            <TableCell align="center">{item2.id}</TableCell>
                            <TableCell align="center">{item2.edges?.typetreatment?.typetreatment}</TableCell>
                            <TableCell align="center">{item2.edges?.patientrecord?.name}</TableCell>   
                            <TableCell align="center">{item.amount}</TableCell>
                            <TableCell align="center">{item.date}</TableCell>                                              
                          </TableRow>
                        ))))}
                      </TableBody>
                    </Table>
                    
                  </TableContainer>
                  <br/>
                    <Button
                        variant = "contained"
                        href = "/createbill"
                        color = "primary"
                    >
                        ย้อนกลับ
                    </Button>
                </Typography>
        </Content>
      </Page>

  );
};

export default BillTable;