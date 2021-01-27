import React, { useEffect, FC } from 'react';

import { Content, Header, Page, pageTheme, } from '@backstage/core';

import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Grid, MenuItem, Button, TextField, Select, Typography } from '@material-ui/core';
import SaveIcon from '@material-ui/icons/Save';

import { Alert } from '@material-ui/lab';
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
import { EntTreatment } from '../../api';
import { EntUser } from '../../api/models/EntUser';
import { Cookies } from 'react-cookie/cjs';//cookie
import Swal from 'sweetalert2'; // alert
import BillTable from '../tableBill';

  const cookies = new Cookies();
  const FINID = cookies.get('Fin'); 
  const Name = cookies.get('Name');
  const Img = cookies.get('Img');

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
      minWidth: 550,
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
  payer?: string;
  payercontact?: string;
  paytype?: number;
  unpaybill?: number;
}

const CreateBill: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(Boolean);
  const [loading, setLoading] = React.useState(true);
  
  const [paytypes, setPaytypes] = React.useState<EntPaytype[]>([]);
  const [financiers, setFinanciers] = React.useState<Partial<EntFinancier>>();

  const [unpaybills, setUnpaybills] = React.useState<EntUnpaybill[]>([]);
  const [treatments, setTreatment] = React.useState<EntTreatment[]>([]);
  
  const [bill, setBill] = React.useState<Partial<Bill>>({});

  const [patientname, setPatient] = React.useState(String);
  const [treatmentid, setTreatmentID] = React.useState(Number);
  const [upbid, setunpaybill] = React.useState(Number);

  const [Users, setUsers] = React.useState<Partial<EntUser>>();
  const [nameError, setnameError] = React.useState('');
  const [phoneError, setphoneError] = React.useState('');
  const [amountError, setamountError] = React.useState('');
   // alert setting
 const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 5000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});
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
    const getTreatment = async () => {
      const res = await http.listTreatment({ offset: 0 });
      setLoading(false);
      setTreatment(res);
    };

    const getImg = async () => {
      const res = await http.getUser({ id: Number(Img) });
      setLoading(false);
      setUsers(res);
    };


    getFinancier();
    getUnpaybill();
    getPaytype();
    getTreatment();
    getImg();
  }, [loading]);

  const refreshPage = () => {
    window.location.reload();
  }
  const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown; id?:string}>) => {
    const name = event.target.name as keyof typeof CreateBill;
    const { value } = event.target;
    const id = event.target.id as string
    const validateValue = value as string
    checkPattern(id, validateValue)
    setBill({ ...bill, [name]: value,unpaybill:upbid,financier:financiers?.id });
    console.log(bill);
  };
  const validatename = (val: string) => {
    return val.charCodeAt(0) >= 65 && val.charCodeAt(0) <= 90 ? true : false;
  }
  const validatetelphone = (val: string) => {
    return val.length == 10 && val.charCodeAt(0) == 48 ? true : false;
  }

  const validateamount = (val: string) => {
    return val.charCodeAt(0) != 48 ? true : false ;
  }

  const checkPattern  = (id: string, value: string) => {
    switch(id) {
      case 'payer':
        validatename(value) ? setnameError('') : setnameError('ชื่อต้องขึ้นต้นด้วยอักษรพิมพ์ใหญ่');
        return;
      case 'payercontact':
        validatetelphone(value) ? setphoneError('') : setphoneError('เบอร์โทรศัพท์ต้องประกอบด้วย 10 ตัวเลข และเริ่มต้นด้วย 0');
        return;
      case 'amount':
        validateamount(value) ? setamountError('') : setamountError('ค่ารักษาต้องเป็นตัวเลข และไม่สามารถเริ่มต้นด้วยเลข 0')
        return;
      default:
        return;
    }
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }
  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'Payer':
        alertMessage("error","ชื่อต้องขึ้นต้นด้วยอักษรพิมพ์ใหญ่");
        return;
      case 'Payercontact':
        alertMessage("error","เบอร์โทรศัพท์ต้องประกอบด้วย 10 ตัวเลข และเริ่มต้นด้วย 0");
        return;
      case 'Amount':
        alertMessage("error","ค่ารักษาต้องเป็นตัวเลข");
        return;
      default:
        alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
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
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
          const upb = {
            id: bill.unpaybill,
          }
          await http.updateUnpaybill({ id: bill.unpaybill as number, unpaybill: upb });
        } else {
          checkCaseSaveError(data.error.Name);

        }
      });
  };

  return (
    
      <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`Financial System`}>

        <Avatar alt="Remy Sharp" src={Users?.images as string} />
        <div style={{ marginLeft: 10 }}>{Name}</div>
      </Header>
        <Content>
          <Grid container spacing={3} >
            <Grid container item xs={12} sm={12}  >
              <Grid item xs={4}>
{/* ********************************* Create Bill Detail ***************************************** */}
                <Paper >
                  <Typography align="center">
                      <Typography align="center" variant="h3">
                        <br /> ใบเสร็จรับเงิน
                      </Typography>
                        <Typography align="center" variant="subtitle1">
                          <br />เลขที่การรักษา : {treatmentid}
                          <br />ผู้ป่วย<br />
                        <TextField
                          disabled
                          label={patientname}
                          size="small"
                          />
                        </Typography>

                        <Typography align="center" variant="subtitle1">
                          <br/>รูปแบบการชำระ<br /> 
                            <Select
                            id = "Paytype"
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
                            <br/>
                          <br/>ค่ารักษา<br />
                            <TextField
                            name = "amount"
                            id = "Amount"
                            className={classes.formControl}
                            value={bill.amount}
                            size = "small"
                            onChange={handleChange} />
                          <br/><br/>ผู้ชำระค่ารักษา<br />
                            <TextField
                            name = "payer"
                            id = "Payer"
                            className={classes.formControl}
                            value={bill.payer}
                            size = "small"
                            onChange={handleChange} />
                          <br/>
                          <br/>เบอร์โทรติดต่อ<br />
                            <TextField
                            name = "payercontact"
                            id = "Payercontact"
                            className={classes.formControl}
                            value={bill.payercontact}
                            size = "small"
                            onChange={handleChange} />
                          <br/>
                          <br/> พนักงานการเงิน : {financiers?.name}
                          <br/>
                          <br/>
                          <Button
                            onClick={() => {
                            Create_Bill();
                            }}
                            startIcon={<SaveIcon />}
                            variant="contained"
                            color="primary"
                          >
                             บันทึกใบเสร็จ
                          </Button>
                          &emsp;
                          <Button
                          variant = "contained"
                          color ="inherit"
                          href = "/BillSearch"
                          >
                            ประวัติการรับเงิน
                          </Button>
                        </Typography>
                  </Typography>
                  <br />
                </Paper>
{/* ************************** Table Show Unpaybil and Treatment Detail **************************  */}
              </Grid>
              <Grid item xs={8}>
                <Paper>
                  <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                      <TableHead>
                        <TableRow>
                          <TableCell align="center" >เลขที่การรักษา</TableCell>
                          <TableCell align="center">ผู้รับการรักษา</TableCell>
                          <TableCell align="center">ประเภทการรักษา</TableCell>
                          <TableCell align="center">เรียกชำระเงิน</TableCell>
                        </TableRow>
                      </TableHead>
                      <TableBody>
                        {unpaybills.map(item => (treatments.filter(t => t.id === item.edges?.edgesOfTreatment?.id).map(item2 => (
                          <TableRow key={item.id}>
                            <TableCell align="center">{item2.id}</TableCell>
                            <TableCell align="center">{item2.edges?.edgesOfPatientrecord?.name}</TableCell>
                            <TableCell align="center">{item2.edges?.edgesOfTypetreatment?.typetreatment}</TableCell>
                            <TableCell align="center">
                              <Button
                                onClick={() => {
                                  setunpaybill(item.id as number);   
                                  setPatient(item2.edges?.edgesOfPatientrecord?.name as string);
                                  setTreatmentID(item2.id as number);
                                }}
                                variant="outlined"
                                color="primary"
                              >
                                ชำระเงิน
                              </Button>
                            </TableCell>
                          </TableRow>
                        ))))}
                      </TableBody>
                    </Table>
                  </TableContainer>
                </Paper>
              </Grid>
            </Grid>
          </Grid>
        </Content>
      </Page>

  );
};

export default CreateBill;