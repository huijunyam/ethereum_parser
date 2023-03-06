package subscription

import (
	"testing"
)

func TestAttach(t *testing.T) {
	type args struct {
		name             string
		input            Observer
		existingObserver []Observer
		hasErr           bool
		want             []Observer
	}

	observer := NewAddressObserver("unique1")
	observer2 := NewAddressObserver("unique2")
	tests := []args{
		{
			name:             "add new observer successfully",
			input:            observer,
			existingObserver: make([]Observer, 0),
			hasErr:           false,
			want:             []Observer{observer},
		},
		{
			name:             "observer exist",
			input:            observer,
			existingObserver: []Observer{observer},
			hasErr:           true,
			want:             []Observer{observer},
		},
		{
			name:             "add more observer to existing list",
			input:            observer,
			existingObserver: []Observer{observer2},
			hasErr:           false,
			want:             []Observer{observer2, observer},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add := "0xef1c6e67703c7bd7107eed8303fbe6ec2554bf6b"
			publisher := NewAddressMonitor(add)
			if len(tt.existingObserver) > 0 {
				for _, o := range tt.existingObserver {
					_, err := publisher.Attach(o)
					if err != nil {
						t.Errorf("adding existing observer err")
					}
				}
			}
			_, err := publisher.Attach(tt.input)
			if tt.hasErr && err == nil {
				t.Errorf("Expected has error, but got no error")
			}
			if !tt.hasErr {
				for i, v := range tt.want {
					if v.GetUser() != publisher.observers[i].GetUser() {
						t.Errorf("Expected '%s', but got '%s'", v.GetUser(), publisher.observers[i].GetUser())
					}
				}
			}
		})
	}
}

func TestDetach(t *testing.T) {
	type args struct {
		name             string
		input            Observer
		existingObserver []Observer
		hasErr           bool
		want             []Observer
	}

	observer := NewAddressObserver("unique1")
	observer2 := NewAddressObserver("unique2")
	observer3 := NewAddressObserver("unique3")
	tests := []args{
		{
			name:             "delete observer successfully",
			input:            observer,
			existingObserver: []Observer{observer},
			hasErr:           false,
			want:             []Observer{},
		},
		{
			name:             "observer not exist",
			input:            observer,
			existingObserver: []Observer{observer2},
			hasErr:           true,
			want:             []Observer{observer2},
		},
		{
			name:             "delete observer in list with more observers",
			input:            observer,
			existingObserver: []Observer{observer3, observer, observer2},
			hasErr:           false,
			want:             []Observer{observer3, observer2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			add := "0xef1c6e67703c7bd7107eed8303fbe6ec2554bf6b"
			publisher := NewAddressMonitor(add)
			if len(tt.existingObserver) > 0 {
				for _, o := range tt.existingObserver {
					_, err := publisher.Attach(o)
					if err != nil {
						t.Errorf("adding existing observer err")
					}
				}
			}
			_, err := publisher.Detach(tt.input)
			if tt.hasErr && err == nil {
				t.Errorf("Expected has error, but got no error")
			}
			if !tt.hasErr {
				for i, v := range tt.want {
					if v.GetUser() != publisher.observers[i].GetUser() {
						t.Errorf("Expected '%s', but got '%s'", v.GetUser(), publisher.observers[i].GetUser())
					}
				}
			}
		})
	}
}
