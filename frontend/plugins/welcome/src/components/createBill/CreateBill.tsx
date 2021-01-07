import React, { useEffect ,FC } from 'react';
import { Content, Header, Page, pageTheme, } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {Grid,MenuItem,Button,TextField,FormControl,Select, Typography} from '@material-ui/core';

import { Alert } from '@material-ui/lab';
import Paper from '@material-ui/core/Paper';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';


import { DefaultApi } from'../../api/apis';
import { EntPaytype } from '../../api/models/EntPaytype';
import { EntUnpaybill } from '../../api/models/EntUnpaybill';
import { EntFinancier } from '../../api/models/EntFinancier';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      button: {
        display: 'block',
        marginTop: theme.spacing(2),
      },
      formControl: {
          width: 400,
        },
        selectEmpty: {
          marginTop: theme.spacing(1),
        },
        paper: {
            padding: theme.spacing(2),
            textAlign: 'center',
            color: theme.palette.text.secondary,
          },
        table: {
            minWidth: 650,
          },
    }),
  );
const CreateBill: FC<{}> = () => {
    const classes = useStyles();
    const http = new DefaultApi();

    const [status, setStatus] = React.useState(false);
    const [alert, setAlert] = React.useState(true);
    const [loading, setLoading] = React.useState(true);

    const [paytypes, setPaytypes] = React.useState<EntPaytype[]>([]);
    const [financiers, setFinanciers] = React.useState<EntFinancier[]>([]);
    const [unpaybills, setUnpaybills] = React.useState<EntUnpaybill[]>([]);

    const [amounts, setamount] = React.useState(String);
    const [datetime, setDatetime] = React.useState(String);
    const [paytypeid, setpaytypeId] = React.useState(Number);
    const [financierid, setfinancierId] = React.useState(Number);
    const [unpayid, setunpayId] = React.useState(Number);

    const refreshPage = ()=>{
      window.location.reload();
    }

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
            const res = await http.listFinancier();
            setLoading(false);
            setFinanciers(res);
        };
        getFinancier();
        getUnpaybill();
        getPaytype();
    }, [loading]);

    const PaytypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setpaytypeId(event.target.value as number);
      };
    const FinancierhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setfinancierId(event.target.value as number);
      };
    const AmounthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setamount(event.target.value as string);
      };
    const handleDatetimeChange = (event: any) => {
        setDatetime(event.target.value as string);
      };
      const CreatePayment = async () => {
          const b = {
            amount: amounts,
            date: datetime + ":00+07:00",
            financier: financierid,
            paytype: paytypeid,
            unpaybill: unpayid,
          };
          const upb = {
              id: unpayid,
          }
          console.log(b);
          await http.updateUnpaybill({id:unpayid,unpaybill:upb});
          const res: any = await http.createBill({ bill : b });
          setStatus(true);
          if (res.id != '') {
            setAlert(true);
          } else {
            setAlert(false);
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
           <Alert severity="warning" style={{ marginTop: 20 }}>
             มีข้อผิดพลาด โปรดลองอีกครั้ง
           </Alert>
         )}
     </div>
    ):null}


    <Page theme={pageTheme.home}>
        <Header title={`Financial Department`}></Header>
      <Content>
        <Grid container spacing = {3} >
          <Grid container item xs = {12} sm = {12}  >
              <Grid item xs = {4}>
                  <Paper >
                <Typography align ="center">
                    <Typography align = "center" variant = "h3">
                      <br/>----  Create Bill  ----
                    </Typography>
                    <Typography align = "center" variant = "h6">
                        <br/>เลขที่การรักษา
                        </Typography>
                        <Typography align = "center" variant = "subtitle2">
                        <br/>เลขที่การรักษาที่เลือกชำระ {unpayid}
                        </Typography>
                  <FormControl className={classes.formControl}>
                        <Typography align = "center" variant = "h6">
                        <br/>รูปแบบการชำระ
                        </Typography>
                        <Select
                        name="paytype"
                        value={paytypeid}
                        onChange={PaytypehandleChange}
                        >
                        {paytypes.map(item => {
                        return (
                        <MenuItem value={item.id}>
                         {item.paytype}
                        </MenuItem>
                        );
                        })}
                        </Select>
                </FormControl>
                <Typography align = "center" variant = "h6">
                     <br/>ค่ารักษา<br/>
                <TextField 
                    className={classes.formControl}
                    value={amounts}
                    onChange={AmounthandleChange}/>
                </Typography>
                <FormControl className={classes.formControl}>
                        <Typography align = "center"variant = "h6">
                        <br/>พนักงานการเงิน
                        </Typography>
                        <Select
                        name="financier"
                        value={financierid}
                        onChange={FinancierhandleChange}
                        >
                        {financiers.map(item => {
                        return (
                        <MenuItem value={item.id}>
                         {item.name}
                        </MenuItem>
                        );
                        })}
                        </Select>
                </FormControl>
                <br/>
                <Typography align = "center"variant = "h6">
                        <br/>วันเวลาที่ชำระ
                        </Typography>
                <TextField
                    className={classes.formControl}
                    id="datetime"
                    type="datetime-local"
                    value={datetime}
                    onChange={handleDatetimeChange}
                    InputLabelProps={{
                     shrink: true,
                     }}
                 />
                <Typography align ="center">
                <br/>
                <Button
                         onClick={() => {
                            CreatePayment();
                            refreshPage();
                          }}
                          
                          variant="contained"
                          color="primary"
                        >
                         Save   
                </Button>
                </Typography> 
            </Typography>
            <br/>
            </Paper>
            <Paper>
            </Paper>
              </Grid>
              <Grid item xs = {8}>
                  <Paper>
                  <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
         <TableCell align="center" >เลขที่การรักษา</TableCell>
         <TableCell align="center">ผู้รับการรักษา</TableCell>
         <TableCell align="center">รูปแบบการรักษา</TableCell>
         <TableCell align="center">เรียกชำระเงิน</TableCell>
         </TableRow>
         </TableHead>
         <TableBody>
                 {unpaybills.filter(b=>b.status === "Unpay").map(item =>(
                <TableRow key={item.id}>
                    <TableCell align="center">{item.edges?.treatment?.id}</TableCell>
                    <TableCell align="center">{item.edges?.treatment?.edges?.patientrecord?.name}</TableCell>
                    <TableCell align="center">{item.edges?.treatment?.edges?.typetreatment?.typetreatment}</TableCell> 
                    <TableCell align="center">
                        <Button
                         onClick={() => {
                            setunpayId(item.id as number);
                          }}
                          variant="outlined"
                          color= "primary"
                        >
                            ชำระเงิน
                        </Button>
                    </TableCell>
                </TableRow>
            ))}
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