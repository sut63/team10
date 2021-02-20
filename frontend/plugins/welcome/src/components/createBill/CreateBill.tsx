import React, { useEffect, FC } from 'react';

import { Content, Header, Page, pageTheme, } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Grid, MenuItem, Button, TextField, Select, Typography } from '@material-ui/core';
import { Alert, AlertTitle } from '@material-ui/lab';
import Paper from '@material-ui/core/Paper';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { Avatar } from '@material-ui/core';


import { DefaultApi } from '../../api/apis';
import { EntPaytype } from '../../api/models/EntPaytype';
import { EntUnpaybill } from '../../api/models/EntUnpaybill';
import { EntFinancier } from '../../api/models/EntFinancier';
import { EntUser } from '../../api/models/EntUser';

import { Cookies } from 'react-cookie/cjs';//cookie



const cookies = new Cookies();
const FINID = cookies.get('Fin');
const Name = cookies.get('Name');
const Img = cookies.get('Img');

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    paper:{
      minWidth:200,
      height:600,
      padding: theme.spacing(2),
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    formControl: {
      width: 170,
      marginTop: theme.spacing(1),
      marginBottom:theme.spacing(2),
      fontSize:3,
    },
    table: {
      minWidth: 550,
    },
    container: {
      maxHeight: 600,
      color: theme.palette.text.secondary,
    },

  }),
);



// header css
const HeaderCustom = {
  minHeight: '50px',
};
interface Bill {
  amount?: string;
  financier?: number;
  note?: string;
  payercontact?: string;
  paytype?: number;
  unpaybill?: number;
}

const CreateBill: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [loading, setLoading] = React.useState(true);
  const [error, setErr] = React.useState(false);
  const [status, setStatus] = React.useState(false);
  const [msg, setAlert] = React.useState(String);
  const [title, setAlerttitle] = React.useState(String);

  const [paytypes, setPaytypes] = React.useState<EntPaytype[]>([]);
  const [financiers, setFinanciers] = React.useState<Partial<EntFinancier>>();

  const [unpaybills, setUnpaybills] = React.useState<EntUnpaybill[]>([]);
  const [bill, setBill] = React.useState<Partial<Bill>>({});

  const [patientname, setPatient] = React.useState(String);
  const [treatmentid, setTreatmentID] = React.useState(Number);
  const [upbid, setunpaybill] = React.useState(Number);
  const [Users, setUsers] = React.useState<Partial<EntUser>>();
 
  useEffect(() => {

    const getPaytype = async () => {
      const res = await http.listPaytype();
      setLoading(false);
      setPaytypes(res);
    };
    const getUnpaybill = async () => {
      const res = await http.listUnpaybill();
      setLoading(false);
      setUnpaybills(res);
    };
    const getFinancier = async () => {
      const res = await http.getFinancier({ id: Number(FINID) });
      setLoading(false);
      setFinanciers(res);
    };

    const getImg = async () => {
      const res = await http.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };


    getFinancier();
    getUnpaybill();
    getPaytype();
    getImg();
  }, [loading]);

  const refreshPage = () => {
    window.location.reload();
  }
  const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown;}>) => {
    const name = event.target.name as keyof typeof CreateBill;
    const { value } = event.target;
    setBill({ ...bill, [name]: value, unpaybill: upbid, financier: financiers?.id });
    console.log(bill);
  };

  const checkCaseSaveError = (field:string) => {
    switch (field) {
      
      case 'Payercontact':
        setAlerttitle("เบอร์โทรศัพท์ไม่ถูกต้อง")
        setAlert("กรุณากรอกเป็นตัวเลข 0-9 และขึ้นต้นด้วยเลข 0")
        return;
      case 'Amount':
        setAlerttitle("ค่ารักษาไม่ถูกต้อง")
        setAlert("กรุณากรอกเป็นตัวเลขและมีค่าไม่น้อยกว่า 0")
        return;
      case 'Note':
        setAlerttitle("หมายเหตุไม่ถูกต้อง")
        setAlert("ไม่สามารถกรอกสัญลักษณ์ในหมายเหตุ")
        return;
      default:
        setAlerttitle("บันทึกข้อมูลไม่สำเร็จ")
        setAlert("กรุณาตรวจสอบข้อมูลอีกครั้ง")

        return;
    }
  }

  const Create_Bill = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/bills';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(bill),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(async data => {
        console.log(data);
        if (data.status === true) {
          const upb = {
            id: bill.unpaybill,
          }
          await http.updateUnpaybill({ id: bill.unpaybill as number, unpaybill: upb });
          
          setStatus(true);
          setErr(false);
          setTimeout(() => {
            setStatus(false);
            refreshPage();
          }, 5000);
        } else {
          checkCaseSaveError(data.error.Name);
          setStatus(true);
          setErr(true);
          setTimeout(() => {
            setStatus(false);
          }, 5000);

        }
      });
  };

  return (

    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ส่วนการเงิน`}>
        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
      <Content>
        <Grid container spacing={3} >
          <Grid container item xs={12} sm={12}  >
            <Grid item xs={4}>
{/* ************************************************ Create Bill Detail ******************************************************* */}
              <Paper className={classes.paper}>
        {status ? (
            <div>
              {error ? (
                <Alert severity="warning">
                  <AlertTitle>{title}</AlertTitle>
                {msg}
                </Alert>
              ) : (
                  <Alert severity="success" style={{ marginTop: 20 }}>
                    <AlertTitle>บันทึกข้อมูลสำเร็จ</AlertTitle>
                  </Alert>
                )}
                <br/>
            </div>
            
          ) : null}
                <Typography align="center">
                  <Typography align="center" variant="h5">
                     ใบเสร็จรับเงิน
                      </Typography>
                  <Typography align="center" variant="subtitle2">
                    <br />เลขที่การรักษา : {treatmentid}
                    <br />ผู้ป่วย : {patientname} <br />
                  </Typography>
                  <Typography align="center" variant="body2">
                    รูปแบบการชำระ<br />
                    <Select
                      id="Paytype"
                      name="paytype"
                      value={bill.paytype}
                      className={classes.formControl}
                      onChange={handleChange}
                    >
                      {paytypes.map(item => {
                        return (
                          <MenuItem value={item.id}>{item.paytype}</MenuItem>
                        );
                      })}
                    </Select>
                    <br />ค่ารักษา<br />
                    <TextField
                      name="amount"
                      id="Amount"
                      className={classes.formControl}
                      value={bill.amount}
                      size="small"
                      onChange={handleChange} />
                   
                    <br />เบอร์โทรติดต่อ<br />
                    <TextField
                      name="payercontact"
                      id="Payercontact"
                      className={classes.formControl}
                      value={bill.payercontact}
                      size="small"
                      onChange={handleChange} />
                    <br />หมายเหตุ<br />
                    <TextField
                      name="note"
                      id="Note"
                      className={classes.formControl}
                      value={bill.note}
                      size="small"
                      onChange={handleChange} />
                    <br />พนักงานการเงิน : {financiers?.name}<br/><br/>

                    <Button
                      onClick={() => {
                        Create_Bill();
                      }}
                      variant="contained"
                      color="primary"
                      style={{backgroundColor: "#00af91",width:150,padding: '6px 12px',}}
                    >
                      บันทึกใบเสร็จ
                          </Button>&emsp;
                          <Button
                      variant="contained"
                      href="/BillSearch"
                      style={{width: 150,padding: '6px 12px',}}
                    >
                      ประวัติการรับเงิน
                          </Button>
                  </Typography>
                </Typography>
                <br />
              </Paper>
            </Grid>
{/* *********************************** Table Show Unpaybil and Treatment Detail ******************************************  */}
            <Grid item xs={8}>
                <TableContainer component={Paper} className={classes.container}>
                  <Table className={classes.table} aria-label="simple table">
                    <TableHead>
                      <TableRow>
                        <TableCell align="center" >เลขที่การรักษา</TableCell>
                        <TableCell align="center">ผู้รับการรักษา</TableCell>
                        <TableCell align="center">ยาที่จ่าย</TableCell>
                        <TableCell align="center">เรียกชำระเงิน</TableCell>
                      </TableRow>
                    </TableHead>
                    <TableBody>
                      {unpaybills.map(item => (
                        <TableRow key={item.id}>
                          <TableCell align="center">{item.edges?.edgesOfTreatment?.id}</TableCell>
                          <TableCell align="center">{item.edges?.edgesOfTreatment?.edges?.edgesOfPatientrecord?.name}</TableCell>
                          <TableCell align="center">{item.edges?.edgesOfTreatment?.medicine}</TableCell>
                          <TableCell align="center">
                            <Button
                              onClick={() => {
                                setunpaybill(item.id as number);
                                setPatient(item.edges?.edgesOfTreatment?.edges?.edgesOfPatientrecord?.name as string);
                                setTreatmentID(item.edges?.edgesOfTreatment?.id as number);
                              }}
                              style={{backgroundColor: "#276678"}}
                              variant = 'contained'
                              color="primary"
                            >
                              ชำระเงิน
                              </Button>
                          </TableCell>
                        </TableRow>
                      ))}
                    </TableBody>
                  </Table>
                </TableContainer>
            </Grid>
          </Grid>
        </Grid>
      </Content>
    </Page>

  );
};

export default CreateBill;