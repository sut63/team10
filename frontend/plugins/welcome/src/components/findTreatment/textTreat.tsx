import React ,{useState,useEffect}from 'react';
import Backdrop from '@material-ui/core/Backdrop';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntTreatment,EntDoctor } from '../../api/models';
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

export default function Textinfo( id : any) {
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
  const [doctors, setDoctors] = useState<EntDoctor[]>(Array);
  const [treatments, setTreatments] = useState<EntTreatment[]>(Array);
  
  const getDoctors = async () => {
    const res = await http.listDoctor({ offset: 0 });
    setLoading(false);
    setDoctors(res);
  };
  const getTreatments = async () => {
    const res = await http.listTreatment({offset : 0});
    setLoading(false);
    setTreatments(res);
  };

  useEffect(() => {
    getDoctors();
    getTreatments();
  
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
          {treatments.filter(t => t.id === id.id).map(item => (doctors.filter(d => d.id === item.edges?.edgesOfDoctor?.id).map(item2 =>  ( 
                    <Grid>
                    <Typography variant = 'h6'>
                        รายละเอียด<br/>
                    </Typography>
                    <Typography variant = 'subtitle1'> 
                        เลขที่บันทึกการรักษา : &emsp;{item.id}<br/>
                        แพทย์ผู้ดูแล :&emsp;{item2.edges?.edgesOfDoctorinfo?.doctorname} {item2.edges?.edgesOfDoctorinfo?.doctorsurname}<br/>
                        ผู้ป่วย : &emsp;{item.edges?.edgesOfPatientrecord?.name}<br/>
                        อาการ : &emsp;{item.symptom}<br/>
                        การรักษา : &emsp;{item.treat}<br/>
                        ยาที่จ่าย : &emsp;{item.medicine}<br/>
                        วันเวลาที่รักษา : &emsp;{item.datetreat}<br/>                  
                    </Typography>
                    </Grid>
                       ))))}
          </div>
        </Fade>
      </Modal>
    </div>
  );
}