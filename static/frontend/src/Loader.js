import { Suspense } from 'react';
import LinearProgress from '@mui/material/LinearProgress';
import { styled } from '@mui/material/styles';

// styles
const LoaderWrapper = styled('div')({
    position: 'fixed',
    top: 0,
    left: 0,
    zIndex: 1301,
    width: '100%'
});

const ProgressLoader = () => (
    <LoaderWrapper>
        <LinearProgress color="primary" />
    </LoaderWrapper>
);

const Loader = (Component) => (props) =>
    (
        <Suspense fallback={<ProgressLoader />}>
            <Component {...props} />
        </Suspense>
    );

export default Loader;

