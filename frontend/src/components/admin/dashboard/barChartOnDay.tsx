import React, { useEffect, useRef, useState } from 'react';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
  ChartEvent,
  ActiveElement,
} from 'chart.js';
import { Bar, getElementAtEvent } from 'react-chartjs-2';
import { ChartOnDayData } from '.';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

type Props = {
  title: string;
  chartData: ChartOnDayData[];
    // eslint-disable-next-line no-unused-vars
  handleChooseHour: (hour: number) => void
};
const labels = [
  '0h',
  '1h',
  '2h',
  '3h',
  '4h',
  '5h',
  '6h',
  '7h',
  '8h',
  '9h',
  '10h',
  '11h',
  '12h',
  '13h',
  '14h',
  '15h',
  '16h',
  '17h',
  '18h',
  '19h',
  '20h',
  '21h',
  '22h',
  '23h',
];
const BarChartOnDay: React.FC<Props> = (props: Props) => {
  const [amountOfJobs, setAmountOfJobs] = useState<Array<number>>([]);
  const chartRef = useRef<ChartJS<'bar'>>();
  const options = {
    responsive: true,
    onHover: (event: ChartEvent, elements: ActiveElement[]) => {
      const target = event.native?.target as HTMLElement
      if (target) {
        target.style.cursor = elements[0] ? 'pointer' : 'default';
      }
    },
    plugins: {
      legend: {
        position: 'top' as const,
      },
      title: {
        display: true,
        text: props.title,
      },
      tooltip: {
        callbacks: {
          label: (context: any) => {
            const labels: Array<string> = [];
            const barData = props.chartData[context.dataIndex];
            labels.push(`amount of cronjob run: ${barData.amount_of_jobs}`);
            barData.cronjobs.forEach((value) => {
              labels.push(`cronjob: ${value.name}, run: ${value.times} times`);
            });
            return labels;
          },
        },
      },
    },
  };

  const getAmountOfJobs = () => {
    const array: Array<number> = [];
    props.chartData.forEach((value) => {
      array.push(value.amount_of_jobs);
    });
    setAmountOfJobs(array);
  };
  useEffect(() => {
    getAmountOfJobs();
  }, [props.chartData]);

  const data = {
    labels,
    datasets: [
      {
        label: 'Jobs run',
        data: amountOfJobs,
        backgroundColor: 'rgba(53, 162, 235, 0.5)',
      },
    ],
  };

  const handleOnClick = (
    event: React.MouseEvent<HTMLCanvasElement, MouseEvent>
  ) => {
    if (chartRef.current) {
      if (getElementAtEvent(chartRef.current, event)[0]) {
        const {index} = getElementAtEvent(chartRef.current, event)[0]
        props.handleChooseHour(index)
      }

    }
    
  };
  return (
    <Bar
      options={options}
      data={data}
      ref={chartRef}
      onClick={(event) => handleOnClick(event)}
    />
  );
};

export default BarChartOnDay;
