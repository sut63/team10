import React ,{useState,useEffect}from 'react';
import Backdrop from '@material-ui/core/Backdrop';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntBill } from '../../api/models';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {Grid, Modal,Typography} from '@material-ui/core';
import Fade from '@material-ui/core/Fade';
const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    modal: {
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
    paper: {
      backgroundColor: theme.palette.background.paper,
      border: '2px solid #000',
      boxShadow: theme.shadows[5],
      padding: theme.spacing(2, 4, 3),
    },
  }),
);
export default function MoreInfo( id : any) {
  const classes = useStyles();
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };
  const http = new DefaultApi();
  const [loading, setLoading] = React.useState(true);
  const [bills, setBills] = useState<EntBill[]>(Array);
  const getBills = async () => {
    const res = await http.listBill({})
    setLoading(false);
    setBills(res);
  };

  useEffect(() => {
    getBills();
  }, [loading]);

  return (
    <div>
      <Button variant='outlined' color = 'primary' onClick={handleOpen}>
        รายละเอียดเพิ่มเติม
      </Button>
      <Modal
        aria-labelledby="transition-modal-title"
        aria-describedby="transition-modal-description"
        className={classes.modal}
        open={open}
        onClose={handleClose}
        closeAfterTransition
        BackdropComponent={Backdrop}
      >
        <Fade in={open}>
          <div className={classes.paper}>
          {bills.filter(b => b.id == id.id).map(bill => (
                    <Grid>
                    <Typography variant = 'h4'>
                        ใบเสร็จรับเงิน<br/>
                    </Typography>
                    <Typography variant = 'subtitle1'> 
                        เลขที่ใบเสร็จ : &emsp;{bill.id}<br/>
                        เลขที่บันทึกการรักษา : &emsp;{bill.id}<br/>
                        แพทย์ผู้ดูแล :&emsp;{bill.edges?.edgesOfUnpaybill?.edges?.edgesOfTreatment?.id}<br/>
                        ผู้ป่วย : &emsp;{bill.edges?.edgesOfUnpaybill?.edges?.edgesOfTreatment?.edges?.edgesOfPatientrecord?.name}<br/>
                        อาการ : &emsp;{bill.edges?.edgesOfUnpaybill?.edges?.edgesOfTreatment?.symptom}<br/>
                        การรักษา : &emsp;{bill.edges?.edgesOfUnpaybill?.edges?.edgesOfTreatment?.treat}<br/>
                        ยาที่จ่าย : &emsp;{bill.edges?.edgesOfUnpaybill?.edges?.edgesOfTreatment?.medicine}<br/>
                        ค่ารักษาทั้งหมด : &emsp;{bill.amount}<br/>
                        รูปแบบการชำระ : &emsp;{bill.edges?.edgesOfPaytype?.paytype}&emsp;
                        
                        เบอร์ติดต่อผู้จ่าย : &emsp;{bill.payercontact}<br/>
                        หมายเหตุ : &emsp;{bill.note}<br/>
                        วันเวลาที่จ่าย : &emsp;{bill.date}<br/>
                        ผู้รับเงิน : &emsp;{bill.edges?.edgesOfOfficer?.name}<br/>
                    </Typography>
                    </Grid>
                       ))}
          </div>
        </Fade>
      </Modal>
    </div>
  );
}