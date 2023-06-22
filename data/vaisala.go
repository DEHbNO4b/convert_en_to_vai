package data

type StrokeVai struct {
	Version        string
	Year           string
	Month          string
	Day            string
	Hour           string
	Minutes        string
	Seconds        string
	Nano           string
	Latitude       string
	Longitude      string
	Signal         string
	Flashes        string
	Sensors        string
	Freedom        string
	Ell_angle      string
	Ell_semimajore string
	Ell_semiminore string
	Chi_square     string
	RiseTime       string
	DownTime       string
	MaxRateRise    string
	Cloud          string
	Ind_angle      string
	Ind_sygnal     string
	Ind_tyme       string
}

func (sv *StrokeVai) String() string {
	var s string
	s = sv.Version + "\t" + sv.Year + "\t" + sv.Month + "\t" + sv.Day + "\t" + sv.Hour + "\t" + sv.Minutes + "\t" + sv.Seconds + "\t"
	s = s + sv.Nano + "\t" + sv.Latitude + "\t" + sv.Longitude + "\t" + sv.Signal + "\t" + sv.Flashes + "\t" + sv.Sensors + "\t"
	s = s + sv.Freedom + "\t" + sv.Ell_angle + "\t" + sv.Ell_semimajore + "\t" + sv.Ell_semiminore + "\t" + sv.Chi_square + "\t"
	s = s + sv.RiseTime + "\t" + sv.DownTime + "\t" + sv.MaxRateRise + "\t" + sv.Cloud + "\t" + sv.Ind_angle + "\t" + sv.Ind_sygnal + "\t" + sv.Ind_tyme
	s = s + "\r\n"
	return s
}
func (sv *StrokeVai) Slice() []string {
	var s = []string{sv.Version, sv.Year, sv.Month, sv.Day, sv.Hour, sv.Minutes, sv.Seconds,
		sv.Nano, sv.Latitude, sv.Longitude, sv.Signal, sv.Flashes, sv.Sensors, sv.Freedom, sv.Ell_angle, sv.Ell_semimajore,
		sv.Ell_semiminore, sv.Chi_square, sv.RiseTime, sv.DownTime, sv.MaxRateRise, sv.Cloud, sv.Ind_angle, sv.Ind_sygnal, sv.Ind_tyme}

	return s

}
