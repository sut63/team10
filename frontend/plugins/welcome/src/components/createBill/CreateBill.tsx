import React, { useEffect, FC } from 'react';
import { Content, Header, Page, pageTheme, } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { Grid, MenuItem, Button, TextField, FormControl, Select, Typography } from '@material-ui/core';
import SaveIcon from '@material-ui/icons/Save';

import { Alert } from '@material-ui/lab';
import Paper from '@material-ui/core/Paper';


import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';


import { DefaultApi } from '../../api/apis';
import { EntPaytype } from '../../api/models/EntPaytype';
import { EntUnpaybill } from '../../api/models/EntUnpaybill';
import { EntFinancier } from '../../api/models/EntFinancier';
import { EntTreatment } from '../../api';

import { Cookies } from 'react-cookie/cjs';//cookie

  const cookies = new Cookies();
  const FINID = cookies.get('Fin');

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

  const [financierid, setfinancierId] = React.useState(Number);
  const [amounts, setamount] = React.useState(String);
  const [paytypeid, setpaytypeId] = React.useState(Number);
  const [unpayid, setunpayId] = React.useState(Number);
  const [patientname, setPatient] = React.useState(String);
  const [treatmentid, setTreatmentID] = React.useState(Number);

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
      const res = await http.listTreatment({ limit: 100, offset: 0 });
      setLoading(false);
      setTreatment(res);
    };
    getFinancier();
    getUnpaybill();
    getPaytype();
    getTreatment();
  }, [loading]);

  const refreshPage = () => {
    window.location.reload();
  }

  const PaytypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setpaytypeId(event.target.value as number);
  };
  const AmounthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setamount(event.target.value as string);
  };

  const CreatePayment = async () => {
    const b = {
      amount: amounts,
      financier: financiers?.id,
      paytype: paytypeid,
      unpaybill: unpayid,
    };
    const upb = {
      id: unpayid,
    }
    console.log(b);
    const res: any = await http.createBill({ bill: b });

    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      await http.updateUnpaybill({ id: unpayid, unpaybill: upb });
      refreshPage();
    } else {
      setAlert(false);
      refreshPage();
    }
    setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <div>
      {status ? (
        <div>
          {alert ? (
            <Alert severity="success">
              บันทึกการชำระสำเร็จ
            </Alert>
          ) : (
              <Alert severity="warning" >
                มีข้อผิดพลาด โปรดลองอีกครั้ง
              </Alert>
            )}
        </div>
      ) : null}
      <Page theme={pageTheme.home}>
        <Header title={`Financial Department`}></Header>
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
                            name="paytype"
                            value={paytypeid}
                            className={classes.formControl}
                            onChange={PaytypehandleChange}
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
                            className={classes.formControl}
                            value={amounts}
                            size = "small"
                            onChange={AmounthandleChange} />
                          <br/>
                          <br/> พนักงานการเงิน : {financiers?.name}
                          <br/>
                          <br/>
                          <Button
                            onClick={() => {
                            CreatePayment();
                            }}
                            startIcon={<SaveIcon />}
                            variant="contained"
                            color="primary"
                          >
                             บันทึกใบเสร็จ
                          </Button>
                        </Typography>
                  </Typography>
                  <br />
                </Paper>
{/* ************************** Table Show Unpaybil and Treatment Detial **************************  */}
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
                        {unpaybills.map(item => (treatments.filter(t => t.id === item.edges?.treatment?.id).map(item2 => (
                          <TableRow key={item.id}>
                            <TableCell align="center">{item2.id}</TableCell>
                            <TableCell align="center">{item2.edges?.patientrecord?.name}</TableCell>
                            <TableCell align="center">{item2.edges?.typetreatment?.typetreatment}</TableCell>
                            <TableCell align="center">
                              <Button
                                onClick={() => {
                                  setunpayId(item.id as number);
                                  setPatient(item2.edges?.patientrecord?.name as string);
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
    </div>
  );
};

export default CreateBill;